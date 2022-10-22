<!-- ========================= SCRIPT -->
<script lang="ts">
	import ResultStore from '../stores/result';
	import ApiPostResolve from '../api/post.resolve';
	import * as StaticInstructions from '../stores/static/instructions';
	import SanitizeInput from '../sanitizers/input';
	import RubikComponent from '../rubik/rubik.svelte';
	import type { Rubik } from '../rubik/rubik';
	import RubikConfig from '../rubik/config';
	import Config from '../config';
	import { getRangedRandomNumber } from '../utils/random';
	import { tick } from 'svelte';

	// --------------------------- constantes
	const screen_rows = 7;
	const screen_columns = 8;
	const max_instructions = screen_rows * screen_columns;
	const max_selected_input = max_instructions - 1;

	// --------------------------- rubiks
	let input_rubik: Rubik | undefined = undefined;
	let output_rubik: Rubik | undefined = undefined;

	// --------------------------- states
	let loading = false;
	let rubik_mode = false;
	$: output_mode = $ResultStore.instructions.length > 0;

	// --------------------------- shaking
	let shake_rbk_button = false;
	let shake_rbk_button_clicked = false;
	let shake_ran_button = false;
	let shake_ran_button_clicked = false;
	let shake_resolve_button = false;
	let shake_resolve_button_clicked = false;

	// --------------------------- prompt/selection
	let prompt_id = 0;
	let prompt_period = false;
	let inputs: string[] = [];
	let selected_input = 0;
	let selected_output = 0;

	let instruction_id = 0;
	let last_play_instruction_id = instruction_id;

	$: selected = output_mode ? selected_output : selected_input;
	$: instructions = output_mode ? $ResultStore.instructions : inputs;

	$: end_selected_instruction = output_mode
		? selected_output === $ResultStore.instructions.length
		: selected_input === inputs.length;
	$: last_selected_instruction = output_mode
		? selected_output === $ResultStore.instructions.length - 1
		: selected_input === inputs.length - 1;
	$: first_selected_instruction = output_mode ? selected_output === 0 : selected_input === 0;
	$: input_str = inputs.join(' ');
	$: input_is_full = inputs.length === max_instructions;
	$: instructions_is_empty = instructions.length === 0;
	$: input_is_empty = inputs.length === 0;

	$: selected_group = output_mode
		? last_selected_instruction
			? '4'
			: ($ResultStore.instructions[selected_output].group - 1).toString()
		: '0';

	// --------------------------- can handle
	$: can_handle_resolve = !(output_mode || input_is_empty || loading) && !loading;
	$: can_handle_play = !(last_selected_instruction || end_selected_instruction) && !loading;
	$: can_handle_play_back =
		!first_selected_instruction &&
		!(end_selected_instruction && instructions.length === 1) &&
		!loading;
	$: can_handle_random = !output_mode && !loading;
	$: can_handle_insert = !output_mode && !loading;
	$: can_handle_reset = (output_mode || !instructions_is_empty) && !loading;
	$: can_handle_delete = can_handle_reset && !loading;
	$: can_handle_instruction = !output_mode && !loading;
	$: can_handle_horizontal_right_move =
		!(end_selected_instruction || (last_selected_instruction && output_mode)) && !loading;
	$: can_handle_horizontal_right_super_move = can_handle_horizontal_right_move;
	$: can_handle_vertical_up_move = can_handle_horizontal_left_move;
	$: can_handle_horizontal_left_move = !first_selected_instruction && !loading;
	$: can_handle_horizontal_left_super_move = can_handle_horizontal_left_move;
	$: can_handle_horizontal_left_super_zero_move =
		can_handle_horizontal_left_move || (first_selected_instruction && !loading);
	$: can_handle_vertical_down_move = can_handle_horizontal_right_move;

	// *************************** utils functions
	// move the prompt safely by checking some conditions
	function selectSafe(index: number): number {
		if (index < 0) return 0;
		if (index > max_selected_input) return max_selected_input;
		if (output_mode) {
			if (index > $ResultStore.instructions.length - 1) return $ResultStore.instructions.length - 1;
		} else {
			if (index > inputs.length) return inputs.length;
		}
		return index;
	}

	function mergeTwoInstruction(first: string, second: string): string {
		const base = second[0];

		switch ((first + second).replaceAll(base, 'X')) {
			case 'XX':
			case "X'X'":
				input_rubik?.pushMove(second);
				return base + '2';
			case 'X2X':
				input_rubik?.pushMove(second);
				return base + "'";
			case "X2X'":
				input_rubik?.pushMove(second);
				return base;
			case "X'X":
				input_rubik?.pushMove(second);
				input_rubik?.pushMove(second);
				return base;
			default:
				input_rubik?.pushMove(first, true);
				input_rubik?.pushMove(second);
				return second;
		}
	}

	// *************************** handlers functions
	// resolve the input by calling the api and show the response in the output
	function handleResolve() {
		if (can_handle_resolve) {
			loading = true;
			instruction_id++;
			shake_resolve_button_clicked = true;

			ApiPostResolve(input_str)
				.then(async (result) => {
					loading = false;
					if (!ResultStore.setFromString(result)) {
						alert('api reponse corrupted');
					} else {
						if ($ResultStore.instructions.length > 0) {
							await tick();
							fromInput();
							await tick();
							output_rubik?.pushMove($ResultStore.instructions[0].instruction);
							await tick();
							handleHorizontalSuperMove(true);
						} else {
							alert('the rubik is already resolved despite the mixture');
						}
					}
				})
				.catch((err) => {
					loading = false;
					alert('an error occured from api');
					console.log(err)
				});
		}
	}

	// add a new instruction
	async function handleInstruction(instruction: string) {
		if (can_handle_instruction) {
			instruction_id++;

			if (!output_mode) {
				const selected_input_value = inputs[selected_input];
				if (selected_input_value != undefined && selected_input_value[0] === instruction[0]) {
					// if overwrite same instruction type

					const merged_instruction = mergeTwoInstruction(selected_input_value, instruction);
					inputs[selected_input] = merged_instruction;
					handleHorizontalMove(true);

					return;
				}

				if (!end_selected_instruction) {
					// if overwrite an different instruction type
					input_rubik?.pushMove(inputs[selected_input], true);
					inputs[selected_input] = instruction;
				} else if (inputs.length < max_instructions) {
					// else if write a new instruction
					if (!first_selected_instruction) {
						// if the selected input is not the first

						const previous_input_value = inputs[selected_input - 1];

						if (previous_input_value != undefined && previous_input_value[0] === instruction[0]) {
							// if overwrite same instruction type
							const merged_instruction = mergeTwoInstruction(previous_input_value, instruction);
							inputs[selected_input - 1] = merged_instruction;

							return;
						}
					}

					// add the new instruction
					inputs = [...inputs, instruction];
				}

				// push the animation
				input_rubik?.pushMove(instruction);
				await tick();
				// move to the right
				handleHorizontalMove(true);
			}
		}
	}

	// move to right and left in the screen
	function handleHorizontalMove(direction: boolean, new_instruction = true) {
		if (
			(direction && can_handle_horizontal_right_move) ||
			(!direction && can_handle_horizontal_left_move)
		) {
			if (new_instruction) instruction_id++;
			new_prompt();

			if (output_mode) {
				const initial_selected_output = selected_output;

				if (direction) {
					// right
					selected_output = selectSafe(selected_output + 1);
				} else {
					// left
					selected_output = selectSafe(selected_output - 1);
				}

				if (selected_output != initial_selected_output) {
					if (
						$ResultStore.instructions[selected_output] != undefined &&
						$ResultStore.instructions[initial_selected_output] != undefined
					) {
						if (selected_output < initial_selected_output) {
							output_rubik?.pushMove(
								$ResultStore.instructions[initial_selected_output].instruction,
								true
							);
						} else {
							output_rubik?.pushMove($ResultStore.instructions[selected_output].instruction, false);
						}
					}
				}
			} else {
				const initial_selected_input = selected_input;

				if (direction) {
					// right
					selected_input = selectSafe(selected_input + 1);
				} else {
					// left
					selected_input = selectSafe(selected_input - 1);
				}

				if (selected_input != initial_selected_input) {
					if (inputs[selected_input] != undefined && inputs[initial_selected_input] != undefined) {
						if (selected_input < initial_selected_input) {
							input_rubik?.pushMove(inputs[initial_selected_input], true);
						} else {
							input_rubik?.pushMove(inputs[selected_input], false);
						}
					}
				}
			}
		}
	}

	// move to start and end in the screen
	function handleHorizontalSuperMove(
		direction: boolean,
		zero = false,
		duration = RubikConfig.moves.durations.nitro
	) {
		if (
			(!direction && can_handle_horizontal_left_super_zero_move) ||
			(direction && can_handle_horizontal_right_super_move) ||
			(!direction && can_handle_horizontal_left_super_move)
		) {
			instruction_id++;
			new_prompt();

			if (output_mode) {
				const initial_selected_output = selected_output;

				if (direction) {
					// right
					selected_output = selectSafe($ResultStore.instructions.length - 1);
				} else {
					// left
					selected_output = selectSafe(0);
				}

				if (selected_output != initial_selected_output || zero) {
					if (!direction || zero) {
						// left
						$ResultStore.instructions
							.slice(zero ? 0 : 1, initial_selected_output + 1)
							.reverse()
							.forEach((instruction) =>
								output_rubik?.pushMove(instruction.instruction, true, duration)
							);
					} else {
						// right
						$ResultStore.instructions
							.slice(initial_selected_output + 1)
							.forEach((instruction) =>
								output_rubik?.pushMove(instruction.instruction, false, duration)
							);
					}
				}
			} else {
				const initial_selected_input = selected_input;

				if (direction) {
					// right
					selected_input = selectSafe(inputs.length);
				} else {
					// left
					selected_input = selectSafe(0);
				}

				if (selected_input != initial_selected_input || zero) {
					if (!direction || zero) {
						// left
						inputs
							.slice(zero ? 0 : 1, initial_selected_input + 1)
							.reverse()
							.forEach((instruction) => input_rubik?.pushMove(instruction, true, duration));
					} else {
						// right
						inputs
							.slice(initial_selected_input + 1)
							.forEach((instruction) => input_rubik?.pushMove(instruction, false, duration));
					}
				}
			}
		}
	}

	// move to up and down in the screen
	function handleVerticalMove(direction: boolean, duration = RubikConfig.moves.durations.fast) {
		if (
			(direction && can_handle_vertical_up_move) ||
			(!direction && can_handle_vertical_down_move)
		) {
			instruction_id++;
			new_prompt();

			if (output_mode) {
				const initial_selected_output = selected_output;
				const initial_end_selected_instruction = end_selected_instruction;

				if (direction) {
					// up
					selected_output = selectSafe(selected_output - screen_columns);
				} else {
					// down
					selected_output = selectSafe(selected_output + screen_columns);
				}

				if (selected_output != initial_selected_output) {
					if (direction) {
						// up
						$ResultStore.instructions
							.slice(
								selected_output + 1,
								initial_selected_output + (initial_end_selected_instruction ? 0 : 1)
							)
							.reverse()
							.forEach((instruction) =>
								output_rubik?.pushMove(instruction.instruction, true, duration)
							);
					} else {
						// down
						$ResultStore.instructions
							.slice(initial_selected_output + 1, selected_output + 1)
							.forEach((instruction) =>
								output_rubik?.pushMove(instruction.instruction, false, duration)
							);
					}
				}
			} else {
				const initial_selected_input = selected_input;
				const initial_end_selected_instruction = end_selected_instruction;

				if (direction) {
					// up
					selected_input = selectSafe(selected_input - screen_columns);
				} else {
					// down
					selected_input = selectSafe(selected_input + screen_columns);
				}

				if (selected_input != initial_selected_input) {
					if (direction) {
						// up
						inputs
							.slice(
								selected_input + 1,
								initial_selected_input + (initial_end_selected_instruction ? 0 : 1)
							)
							.reverse()
							.forEach((instruction) => input_rubik?.pushMove(instruction, true, duration));
					} else {
						// down
						inputs
							.slice(initial_selected_input + 1, selected_input + 1)
							.forEach((instruction) => input_rubik?.pushMove(instruction, false, duration));
					}
				}
			}
		}
	}

	// play instruction by instruction each corresponding animation
	function handlePlay(direction: boolean, from_instruction_id: number | undefined = undefined) {
		if ((direction && can_handle_play) || (!direction && can_handle_play_back)) {
			if (last_play_instruction_id === instruction_id && from_instruction_id === undefined) {
				// if play status is interupted #1
				instruction_id++;
				return;
			}

			if ((end_selected_instruction && direction) || (first_selected_instruction && !direction)) {
				// if play status is interupted #2
				instruction_id++;
				return;
			}

			const local_instruction_id =
				from_instruction_id != undefined ? from_instruction_id : ++instruction_id;
			last_play_instruction_id = local_instruction_id;

			setTimeout(() => {
				// create a cycle by recursivity
				if (local_instruction_id === instruction_id) {
					handleHorizontalMove(direction, false);
					// call the next move
					handlePlay(direction, local_instruction_id);
				}
			}, RubikConfig.moves.durations.play);
		}
	}

	// reset input and output
	async function handleReset() {
		if (can_handle_reset) {
			instruction_id++;
			const from_output_mode = output_mode;

			if (from_output_mode) {
				// if in output mode
				// reset outputs
				handleHorizontalSuperMove(false, true, RubikConfig.moves.durations.nitro);
				ResultStore.reset();
				await tick();
			}

			// reset input
			handleHorizontalSuperMove(false, true, RubikConfig.moves.durations.nitro);
			if (from_output_mode) toInput();
			inputs = [];

			new_prompt();
		}
	}

	// delete an instruction or reset output
	function handleDelete() {
		if (can_handle_delete) {
			instruction_id++;

			if (output_mode) {
				// if in output mode
				// reset the output result and back to inputs
				handleHorizontalSuperMove(false, true, RubikConfig.moves.durations.nitro);
				toInput();
				ResultStore.reset();
			} else {
				const initial_instruction = inputs[selected_input];

				// delete the current selected instruction
				if (last_selected_instruction) {
					inputs = inputs.slice(0, selected_input);
				} else if (!end_selected_instruction) {
					inputs = [...inputs.slice(0, selected_input), ...inputs.slice(selected_input + 1)];
				} else {
					if (!input_is_empty) {
						input_rubik?.pushMove(inputs[selected_input - 1], true);
						inputs = [...inputs.slice(0, selected_input - 1)];
					}

					handleHorizontalMove(false);
				}

				// get the new current selected instruction
				const superseding_instruction = inputs[selected_input];

				if (initial_instruction != superseding_instruction) {
					// push the correspondig animation
					if (initial_instruction != undefined) {
						input_rubik?.pushMove(initial_instruction, true);
						if (superseding_instruction != undefined) {
							input_rubik?.pushMove(superseding_instruction, false);
						}
					}
				}
			}
		}
	}

	// insert a new default instruction
	function handleInsert() {
		if (can_handle_insert) {
			instruction_id++;

			if (!output_mode) {
				if (input_is_full) {
					// if the input is full
					handleInstruction(StaticInstructions.physical_instructions[0]);
				} else {
					const initial_instruction = inputs[selected_input];

					inputs = [
						...inputs.slice(0, selected_input),
						StaticInstructions.instructions[0],
						...inputs.slice(selected_input)
					];

					const superseding_instruction = inputs[selected_input];

					if (initial_instruction != superseding_instruction) {
						// push the correspondig animation
						if (initial_instruction != undefined) input_rubik?.pushMove(initial_instruction, true);
						if (superseding_instruction != undefined)
							input_rubik?.pushMove(superseding_instruction, false);
					}
				}
			}
		}
	}

	// show or hide rubik visualization
	function handleVisualization() {
		shake_rbk_button_clicked = true;
		rubik_mode = !rubik_mode;
	}

	// generate a new input random set of instruction
	async function handleRandom() {
		if (can_handle_random) {
			instruction_id++;
			shake_ran_button_clicked = true;

			if (!output_mode) {
				// back to the initiale state
				handleHorizontalSuperMove(false, true);

				// get the random number of instruction to generate
				const instruction_nbr = getRangedRandomNumber(1, max_instructions);

				// reset the inputs
				inputs = [];

				// generate the random instruction to inputs
				for (let i = 0; i < instruction_nbr; i++)
					inputs.push(
						StaticInstructions.instructions[
							getRangedRandomNumber(0, StaticInstructions.instructions.length)
						]
					);

				// push the first instruction animation if exists
				if (inputs.length > 0) {
					await tick();
					input_rubik?.pushMove(inputs[0], false, RubikConfig.moves.durations.nitro);
					handleHorizontalSuperMove(true);
				}
			}
		}
	}

	// listen keyboard events
	let HandleKeyDown = (event: any) => {
		switch (event.keyCode) {
			case 32:
				// space
				handleResolve();
				break;
			case 13:
				// enter
				handleInsert();
				break;
			case 8:
				// backspace
				handleDelete();
				break;
			case 46:
				// delete
				handleReset();
				break;
			case 82:
				// r
				handleRandom();
				break;
			case 88:
				// x
				handleVisualization();
				break;
			case 37:
				// left
				handleHorizontalMove(false);
				break;
			case 38:
				// up
				handleVerticalMove(true);
				break;
			case 39:
				// left
				handleHorizontalMove(true);
				break;
			case 38:
				// down
				handleVerticalMove(false);
				break;
			case 80:
				// p
				handlePlay(true);
				break;
			case 66:
				// p
				handlePlay(false);
				break;
			default:
				break;
		}
	};

	// *************************** clipboard functions
	// handle the copy methode of the clipboard
	function handleCopy() {
		if (output_mode) {
			navigator.clipboard.writeText($ResultStore.instructions.join(' '));
		} else {
			navigator.clipboard.writeText(input_str);
		}
	}

	// handle the past methode of the clipboard
	function handlePaste() {
		if (!output_mode) {
			navigator.clipboard.readText().then((cliptext) => {
				try {
					const clip_instructions = SanitizeInput(cliptext);
					inputs = clip_instructions;
				} catch {
					alert('input bad formatted');
				}
			});
		}
	}

	// *************************** prompt functions
	// manage the prompt cycle for create a blink effect
	function prompt_cycle() {
		const current_prompt_id = prompt_id;

		setTimeout(() => {
			if (current_prompt_id === prompt_id) {
				// if the prompt is the same as at the start
				prompt_period = !prompt_period;
				// continue the cycle by recursivity
				prompt_cycle();
			}
		}, 500);
	}

	// restart the prompt by creating a new one
	function new_prompt() {
		// change the prompt id for delte current prompt cycle
		prompt_id += 1;
		// active the new prompt
		prompt_period = true;
		// start a new cycle for the new prompt
		prompt_cycle();
	}

	// *************************** input <> output functions
	// play the input mixture animations for inintialize the output rubik
	function toInput() {
		inputs
			.slice()
			.reverse()
			.forEach((instruction) => output_rubik?.pushMove(instruction, true, 0));
	}

	// play back the input mixture animations for reset the output rubik
	function fromInput() {
		inputs.forEach((instruction) => output_rubik?.pushMove(instruction, false, 0));
	}

	// *************************** shaker functions
	// shake the resolve button randomly
	function resolveShake() {
		setTimeout(() => {
			shake_resolve_button = false;
			setTimeout(() => {
				if (!shake_resolve_button_clicked) {
					if (can_handle_resolve) {
						shake_resolve_button = true;
					}
					resolveShake();
				}
			}, getRangedRandomNumber(Config.shake_animation.min_time_to_start, Config.shake_animation.max_time_to_start));
		}, Config.shake_animation.pause);
	}

	// shake the random button randomly
	function shakeRan() {
		setTimeout(() => {
			shake_ran_button = false;
			setTimeout(() => {
				if (!shake_ran_button_clicked && !shake_resolve_button_clicked) {
					if (can_handle_random) {
						shake_ran_button = true;
					}
					shakeRan();
				} else {
					setTimeout(() => {
						if (!shake_resolve_button_clicked && can_handle_resolve) {
							shake_resolve_button = true;
						}
						resolveShake();
					}, Config.shake_animation.next);
				}
			}, getRangedRandomNumber(Config.shake_animation.min_time_to_start, Config.shake_animation.max_time_to_start));
		}, Config.shake_animation.pause);
	}

	// shake the rubik viewer button randomly
	function shakeRbk() {
		setTimeout(() => {
			shake_rbk_button = false;
			setTimeout(() => {
				if (
					!shake_rbk_button_clicked &&
					!shake_ran_button_clicked &&
					!shake_resolve_button_clicked
				) {
					shake_rbk_button = true;
					shakeRbk();
				} else {
					setTimeout(() => {
						if (!shake_ran_button_clicked && !shake_resolve_button_clicked && can_handle_random) {
							shake_ran_button = true;
						}
						shakeRan();
					}, Config.shake_animation.next);
				}
			}, getRangedRandomNumber(Config.shake_animation.min_time_to_start, Config.shake_animation.max_time_to_start));
		}, Config.shake_animation.pause);
	}

	// shake the shake flow
	function startShake() {
		setTimeout(() => {
			if (!shake_rbk_button_clicked) {
				shake_rbk_button = true;
			}
			shakeRbk();
		}, Config.shake_animation.first);
	}

	// *************************** about functions
	// scroll to the bottom of the window
	function scrollToAbout() {
		const html_element = document.getElementById('html');
		html_element?.scrollTo({ top: html_element.scrollHeight, behavior: 'smooth' });
	}

	startShake();
	new_prompt();
</script>

<!-- ========================= HTML -->
<div class="flow-container">
	<div class="text-container">
		<div class="display">
			<div class="screen" style="opacity: {rubik_mode ? 0.1 : 1};">
				<div class="header">
					<p class="inline-block w-14">
						{loading ? 'loading' : output_mode ? 'output' : 'input'}
					</p>
					<p>
						<span class="w-[18px] text-end">{selected + Number(!end_selected_instruction)}</span
						>/<span class="w-[18px] text-end">{instructions.length}</span>
					</p>
					<p style="opacity: {output_mode ? 1 : 0.3};">G{selected_group}</p>
					<p style="opacity: {output_mode ? 1 : 0.3};" class="w-[66px] text-end">
						{$ResultStore.time} ms
					</p>
				</div>
				<div class="instructions">
					{#if output_mode}
						<!-- OUTPUT -->
						{#each $ResultStore.instructions as result_instruction, index}
							<p
								class:selected-input={index === selected && prompt_period}
								class="screen-instruction"
							>
								{result_instruction.instruction}
							</p>
						{/each}
					{:else}
						<!-- INPUT -->
						{#each inputs as instruction, index}
							<p
								class:selected-input={index === selected && prompt_period}
								class="screen-instruction"
							>
								{instruction}
							</p>
						{/each}
					{/if}
					{#if !output_mode && inputs.length < max_instructions}
						<span class:selected-input={end_selected_instruction && prompt_period} class="input"
							>&nbsp;</span
						>
					{/if}
				</div>
			</div>
			<RubikComponent show={rubik_mode} bind:input_rubik bind:output_rubik {output_mode} />
		</div>
		{#if !rubik_mode}
			<div class="clipboard-container">
				<button on:click={handleCopy}>copy</button>
				{#if !output_mode}
					|
					<button on:click={handlePaste}>paste</button>
				{/if}
			</div>
		{/if}
		<a href="https://github.com/trixky/rubik" target="_blank">
			<h1 class="imprimed-title">
				<spane class="text-red-300">R</spane><spane class="text-green-300">u</spane><spane
					class="text-yellow-300">b</spane
				><spane class="text-blue-300">i</spane><spane class="text-orange-300">k</spane><spane
					class="text-neutral-300">6</spane
				><spane class="text-neutral-300">4</spane>
			</h1>
		</a>
		<div class="physic-button-container">
			<button class="invisible" disabled>?</button>
			<button
				title="reset"
				class="physic-button left-rotation red-button"
				on:click={handleReset}
				disabled={!can_handle_reset}>rst</button
			>
			<button
				title="delete / back to input"
				class="physic-button left-rotation red-button"
				on:click={handleDelete}
				disabled={!can_handle_delete}>del</button
			>
			<button
				title="insert"
				class="physic-button right-rotation"
				on:click={handleInsert}
				disabled={!can_handle_insert}>ins</button
			>
			<button
				title="random"
				class:right-shake={shake_ran_button}
				class="physic-button right-rotation"
				on:click={handleRandom}
				disabled={!can_handle_random}>ran</button
			>
			<button class="invisible" disabled>?</button>
			{#each StaticInstructions.physical_instructions as instruction, index}
				<button
					title={'[' + instruction.toLocaleLowerCase() + '] instruction'}
					class="physic-button instruction-button {index % 6 >= 3
						? 'right-rotation'
						: 'left-rotation'}"
					on:click={() => handleInstruction(instruction)}
					disabled={!can_handle_instruction}>{instruction.toLocaleLowerCase()}</button
				>
			{/each}
			<button
				title="move up"
				class="physic-button left-rotation move-button"
				style="rotate: 180deg;"
				on:click={() => handleVerticalMove(true)}
				disabled={!can_handle_vertical_up_move}>{'v'}</button
			>
			<button
				title="move at the start"
				class="physic-button left-rotation move-button"
				on:click={() => handleHorizontalSuperMove(false)}
				disabled={!can_handle_horizontal_left_super_move}>{'<<'}</button
			>
			<button
				title="move left"
				class="physic-button left-rotation move-button"
				on:click={() => handleHorizontalMove(false)}
				disabled={!can_handle_horizontal_left_move}>{'<'}</button
			>
			<button
				title="move right"
				class="physic-button right-rotation move-button"
				on:click={() => handleHorizontalMove(true)}
				disabled={!can_handle_horizontal_right_move}>{'>'}</button
			>
			<button
				title="move at the end"
				class="physic-button right-rotation move-button"
				on:click={() => handleHorizontalSuperMove(true)}
				disabled={!can_handle_horizontal_right_super_move}>{'>>'}</button
			>
			<button
				title="move down"
				class="physic-button right-rotation move-button"
				on:click={() => handleVerticalMove(false)}
				disabled={!can_handle_vertical_down_move}>{'v'}</button
			>
			<button class="invisible" disabled>?</button>
			<button
				title="rubik viewer"
				class:left-shake={shake_rbk_button}
				class="physic-button left-rotation bottom-button"
				on:click={handleVisualization}>rbk</button
			>
			<button
				title="play back / stop"
				class="physic-button left-rotation move-button"
				on:click={() => handlePlay(false)}
				disabled={!can_handle_play_back}>{'<~'}</button
			>
			<button
				title="play / stop"
				class="physic-button right-rotation move-button"
				on:click={() => handlePlay(true)}
				disabled={!can_handle_play}>{'~>'}</button
			>
			<button
				title="resolve instructions"
				class:right-shake={shake_resolve_button}
				class="physic-button right-rotation bottom-button"
				on:click={handleResolve}
				disabled={!can_handle_resolve}>rsl</button
			>
			<button class="invisible" disabled>?</button>
		</div>
	</div>
	<div class="about">
		<h2 on:click={scrollToAbout}>About</h2>
	</div>
</div>

<svelte:window on:keydown|preventDefault={HandleKeyDown} />

<!-- ========================= CSS -->
<style lang="postcss">
	.imprimed-title {
		@apply absolute -top-[28px] left-0 ml-3 select-none;
		font-family: 'Chistoso';
		font-size: 1.1em;
	}

	.flow-container {
		@apply w-fit m-auto p-8 flex rounded-xl;
	}

	.text-container {
		@apply relative flex flex-col items-center;
	}

	.screen {
		@apply relative m-0 p-3 border-solid border-[1px] border-black rounded-md break-words duration-300;
		font-family: 'Minecraftia';
		box-shadow: inset 0 0 10px rgba(0, 0, 0, 0.322);
		background-color: var(--main-screen-color);
	}

	.header {
		@apply m-0 px-[3px] flex justify-between text-sm;
	}

	.header > p,
	.header > p > span {
		@apply inline-block;
	}

	.instructions {
		@apply grid grid-cols-8 grid-rows-7 m-0 pt-2 w-[240px] h-[200px];
	}

	.physic-button-container {
		@apply grid grid-cols-6 grid-rows-5 gap-2 mt-6 mb-0;
	}

	.physic-button-container > button {
		@apply border-solid border-[1px] border-black rounded-md;
	}

	.right-rotation {
		transform: rotate(var(--main-small-right-rotation));
	}

	.left-rotation {
		transform: rotate(var(--main-small-left-rotation));
	}

	.physic-button {
		@apply border-none px-[7px] py-[8px];
		font-weight: 400;
		background-color: #eee;
		font-size: small;
	}

	.physic-button:not(disabled):not([disabled]):not(:disabled):hover {
		filter: brightness(92%);
		cursor: pointer;
	}

	.physic-button:not(disabled):not([disabled]):not(:disabled):hover:active {
		filter: brightness(80%);
	}

	.physic-button:focus {
		outline: none;
	}

	.instruction-button {
		background-color: azure;
	}

	.move-button {
		background-color: snow;
	}

	.bottom-button {
		background-color: rgb(255, 251, 235); /* cornsilk; */
	}

	.physic-button.right-rotation {
		@apply ml-[2px];
	}

	.physic-button.left-rotation {
		@apply ml-[2px];
		margin-right: 2px;
	}

	.red-button {
		@apply text-red-400;
	}

	.screen-instruction {
		@apply p-1 m-0 text-neutral-900 select-none;
		font-size: 85%;
		font-weight: 500;
	}

	.selected-input {
		@apply bg-neutral-900;
		color: #c9e9c5;
	}

	.clipboard-container {
		@apply absolute -top-[34px] right-2 opacity-0 duration-300 p-2 select-none;
	}

	.clipboard-container:hover,
	.display:hover + .clipboard-container {
		@apply opacity-[15%];
	}

	.clipboard-container > button {
		@apply border-none bg-inherit hover:cursor-pointer;
	}

	button:disabled,
	button[disabled] {
		@apply opacity-50;
		border: 1px solid buttonborder;
		color: currentColor;
	}

	button:disabled.red-button,
	button[disabled].red-button {
		@apply text-red-400;
	}

	button.left-shake {
		animation: left-shake 0.82s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
	}

	button.right-shake {
		animation: right-shake 0.82s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
	}

	.about {
		@apply absolute text-center m-auto w-full left-0 bottom-0 mb-3;
	}

	.about > h2 {
		@apply inline-block opacity-20 hover:opacity-100 transition-all duration-200 hover:cursor-pointer;
		font-size: 1.1em;
	}
</style>
