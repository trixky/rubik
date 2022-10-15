<!-- ========================= SCRIPT -->
<script lang="ts">
	import ResultStore from '../stores/result';
	import ApiPostResolve from '../api/post.resolve';
	import * as StaticInstructions from '../stores/static/instructions';
	import SanitizeInput from '../sanitizers/input';
	import RubikComponent from '../rubik/rubik.svelte';
	import type { Rubik } from '../rubik/rubik';

	const screen_rows = 8;
	const screen_columns = 8;
	const max_instructions = screen_rows * screen_columns;
	const max_selected_input = max_instructions - 1;

	let input_rubik: Rubik | undefined = undefined;
	let output_rubik: Rubik | undefined = undefined;

	$: output_mode = $ResultStore.length > 0;
	let rubik_mode = false;

	let prompt_id = 0;
	let prompt_period = false;
	let inputs: string[] = [];
	let selected_input = 0;
	let selected_output = 0;

	$: selected = output_mode ? selected_output : selected_input;

	$: end_selected = output_mode
		? selected_output === $ResultStore.length
		: selected_input === inputs.length;
	$: last_selected_input = output_mode
		? selected_output === $ResultStore.length - 1
		: selected_input === inputs.length - 1;
	$: input_str = inputs.join(' ');
	$: input_is_full = inputs.length === max_instructions;

	function selectSafe(index: number): number {
		if (index < 0) return 0;
		if (index > max_selected_input) return max_selected_input;
		if (output_mode) {
			if (index > $ResultStore.length - 1) return $ResultStore.length - 1;
		} else {
			if (index > inputs.length) return inputs.length;
		}
		return index;
	}

	function handleResolve() {
		ApiPostResolve(input_str).then((result) => {
			ResultStore.setFromString(result);
		});
	}

	function handleInstruction(instruction: string) {
		if (output_mode) {
			// search
		} else {
			const selected_input_value = inputs[selected_input];
			if (selected_input_value != undefined && selected_input_value[0] === instruction[0]) {
				// If overwrite same instruction type
				const base = instruction[0];

				switch ((selected_input_value + instruction).replaceAll(base, 'X')) {
					case 'XX':
					case "X'X'":
						inputs[selected_input] = base + '2';
						input_rubik?.pushMove(instruction);
						handleHorizontalMove(true, false);
						break;
					case 'X2X':
						inputs[selected_input] = base + "'";
						input_rubik?.pushMove(instruction);
						handleHorizontalMove(true, false);
						break;
					case "X2X'":
						inputs[selected_input] = base;
						input_rubik?.pushMove(instruction);
						handleHorizontalMove(true, false);
						break;
					case "X'X":
						input_rubik?.pushMove(instruction);
						input_rubik?.pushMove(instruction);
						inputs[selected_input] = base;
						handleHorizontalMove(true, false);
						break;
					default:
						input_rubik?.pushMove(inputs[selected_input], true);
						input_rubik?.pushMove(instruction);
						inputs[selected_input] = instruction;
						break;
				}

				return;
			}

			if (!end_selected) {
				// If overwrite an different instruction type
				input_rubik?.pushMove(inputs[selected_input], true);
				inputs[selected_input] = instruction;
			} else if (inputs.length < max_instructions) {
				// If write a new instruction
				inputs = [...inputs, instruction];
			}

			input_rubik?.pushMove(instruction);
			handleHorizontalMove(true, false);
		}
	}

	function handleHorizontalMove(direction: boolean, animation = true) {
		new_prompte();

		if (output_mode) {
			if (direction) {
				// right
				selected_output = selectSafe(selected_output + 1);
			} else {
				// left
				selected_output = selectSafe(selected_output - 1);
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

	function handleHorizontalSuperMove(direction: boolean, zero = false) {
		new_prompte();

		if (output_mode) {
			if (direction) {
				// right
				selected_output = selectSafe($ResultStore.length);
			} else {
				// left
				selected_output = selectSafe(0);
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
						.forEach((instruction) => input_rubik?.pushMove(instruction, true));
				} else {
					// right
					inputs
						.slice(initial_selected_input + 1)
						.forEach((instruction) => input_rubik?.pushMove(instruction, false));
				}
			}
		}
	}

	function handleVerticalMove(direction: boolean) {
		new_prompte();

		if (output_mode) {
			if (direction) {
				// up
				selected_output = selectSafe(selected_output - screen_rows);
			} else {
				// down
				selected_output = selectSafe(selected_output + screen_rows);
			}
		} else {
			const initial_selected_input = selected_input;
			const initial_end_selected = end_selected;

			if (direction) {
				// up
				selected_input = selectSafe(selected_input - screen_rows);
			} else {
				// down
				selected_input = selectSafe(selected_input + screen_rows);
			}

			if (selected_input != initial_selected_input) {
				if (direction) {
					// up
					inputs
						.slice(selected_input + 1, initial_selected_input + (initial_end_selected ? 0 : 1))
						.reverse()
						.forEach((instruction) => input_rubik?.pushMove(instruction, true));
				} else {
					// down
					inputs
						.slice(initial_selected_input + 1, selected_input + 1)
						.forEach((instruction) => input_rubik?.pushMove(instruction, false));
				}
			}
		}
	}

	function handleReset() {
		handleHorizontalSuperMove(false, true);
		selected_input = selectSafe(0);
		selected_output = selectSafe(0);
		inputs = [];
		ResultStore.reset();
		new_prompte();
	}

	function handleInsert() {
		if (!output_mode) {
			if (input_is_full) {
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
					if (initial_instruction != undefined) input_rubik?.pushMove(initial_instruction, true);
					if (superseding_instruction != undefined)
						input_rubik?.pushMove(superseding_instruction, false);
				}
			}
		}
	}

	function handleDelete() {
		if (output_mode) {
			selected_output = 0;
			ResultStore.reset();
		} else {
			const initial_instruction = inputs[selected_input];

			if (last_selected_input) {
				inputs = inputs.slice(0, selected_input);
			} else if (!end_selected) {
				inputs = [...inputs.slice(0, selected_input + 1), ...inputs.slice(selected_input + 2)];
			}

			const superseding_instruction = inputs[selected_input];

			if (initial_instruction != superseding_instruction) {
				if (initial_instruction != undefined) {
					input_rubik?.pushMove(initial_instruction, true);
					if (superseding_instruction != undefined) {
						input_rubik?.pushMove(superseding_instruction, false);
					}
				}
			}
		}
	}

	function prompt_cycle() {
		const current_prompt_id = prompt_id;
		setTimeout(() => {
			if (current_prompt_id === prompt_id) {
				prompt_period = !prompt_period;
				prompt_cycle();
			}
		}, 500);
	}

	function handleDimension() {
		rubik_mode = !rubik_mode;
	}

	function handleRandom() {
		if (!output_mode) {
			handleHorizontalSuperMove(false, true);

			function getRandomNumber(min: number, max: number): number {
				return Math.floor(min + Math.random() * max);
			}

			const instruction_nbr = getRandomNumber(1, max_instructions);

			inputs = [];

			for (let i = 0; i < instruction_nbr; i++)
				inputs.push(
					StaticInstructions.instructions[
						getRandomNumber(0, StaticInstructions.instructions.length)
					]
				);

			if (inputs.length > 0) input_rubik?.pushMove(inputs[0]);
		}
	}

	function handleCopy() {
		if (output_mode) {
			navigator.clipboard.writeText($ResultStore.join(' '));
		} else {
			navigator.clipboard.writeText(input_str);
		}
	}

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

	function new_prompte() {
		prompt_id += 1;
		prompt_period = true;
		prompt_cycle();
	}

	new_prompte();
</script>

<!-- ========================= HTML -->
<div class="flow-container">
	<div class="text-container">
		<div class="screen">
			{#each output_mode ? $ResultStore : inputs as instruction, index}
				<p class:selected-input={index === selected && prompt_period} class="screen-instruction">
					{instruction}
				</p>
			{/each}
			{#if !output_mode && inputs.length < max_instructions}
				<span class:selected-input={end_selected && prompt_period} class="input">&nbsp;</span>
			{/if}
		</div>
		<RubikComponent show={rubik_mode} bind:input_rubik bind:output_rubik {output_mode} />
		<div class="clipboard-container">
			<button on:click={handleCopy}>copy</button>
			{#if !output_mode}
				|
				<button on:click={handlePaste}>paste</button>
			{/if}
		</div>
		<p class="imprimed-title">
			<spane class="text-red-300">R</spane><spane class="text-green-300">u</spane><spane
				class="text-yellow-300">b</spane
			><spane class="text-blue-300">i</spane><spane class="text-orange-300">k</spane><spane
				class="text-neutral-300">6</spane
			><spane class="text-neutral-300">4</spane>
		</p>
		<div class="physic-button-container">
			<button class="physic-button left-rotation text-red-400" on:click={handleReset}>rst</button>
			<button class="physic-button left-rotation text-red-400" on:click={handleDelete}>del</button>
			<button class="physic-button left-rotation" on:click={handleInsert}>ins</button>
			<button class="physic-button right-rotation">grp</button>
			<button class="physic-button right-rotation" on:click={handleRandom}>ran</button>
			<button class="physic-button right-rotation">on</button>
			{#each StaticInstructions.physical_instructions as instruction, index}
				<button
					class="physic-button instruction-button {index % 6 >= 3
						? 'right-rotation'
						: 'left-rotation'}"
					on:click={() => handleInstruction(instruction)}>{instruction.toLocaleLowerCase()}</button
				>
			{/each}
			<button
				class="physic-button left-rotation move-button"
				on:click={() => handleVerticalMove(false)}>{'v'}</button
			>
			<button
				class="physic-button left-rotation move-button"
				on:click={() => handleHorizontalSuperMove(false)}>{'<<'}</button
			>
			<button
				class="physic-button left-rotation move-button"
				on:click={() => handleHorizontalMove(false)}>{'<'}</button
			>
			<button
				class="physic-button right-rotation move-button"
				on:click={() => handleHorizontalMove(true)}>{'>'}</button
			>
			<button
				class="physic-button right-rotation move-button"
				on:click={() => handleHorizontalSuperMove(true)}>{'>>'}</button
			>
			<button
				class="physic-button right-rotation move-button"
				style="rotate: 180deg;"
				on:click={() => handleVerticalMove(true)}>{'v'}</button
			>
			<button class="physic-button left-rotation bottom-button">itr</button>
			<button class="physic-button left-rotation bottom-button">an</button>
			<button class="physic-button left-rotation move-button">{'<~'}</button>
			<button class="physic-button right-rotation move-button">{'~>'}</button>
			<button class="physic-button right-rotation bottom-button" on:click={handleDimension}
				>rbk</button
			>
			<button class="physic-button right-rotation bottom-button" on:click={handleResolve}
				>rsl</button
			>
		</div>
	</div>
</div>

<!-- ========================= CSS -->
<style lang="postcss">
	.imprimed-title {
		@apply absolute -top-[28px] left-2 ml-3 select-none;
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
		@apply grid grid-cols-8 grid-rows-8 m-0 w-[240px] h-[240px] p-3 border-solid border-[1px] border-black rounded-md break-words duration-300;
		box-shadow: inset 0 0 10px rgba(0, 0, 0, 0.322);
		background-color: #c9e9c5;
	}

	.physic-button-container {
		@apply grid grid-cols-6 grid-rows-5 gap-2 mt-5 mb-0;
	}

	.physic-button-container > button {
		@apply border-solid border-[1px] rounded-md;
	}

	.right-rotation {
		transform: rotate(8deg);
	}

	.left-rotation {
		transform: rotate(-8deg);
	}

	.physic-button {
		@apply border-none px-[7px] py-[8px];
		font-weight: 500;
		background-color: #eee;
	}

	.physic-button:hover {
		filter: brightness(92%);
		cursor: pointer;
	}

	.physic-button:hover:active {
		filter: brightness(80%);
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

	.screen-instruction {
		@apply p-1 m-0 text-neutral-900 select-none;
		font-family: 'Minecraftia';
		font-size: 85%;
		font-weight: 500;
	}

	.selected-input {
		@apply bg-neutral-900;
		color: #c9e9c5;
	}

	.clipboard-container {
		@apply absolute -top-[34px] right-3 opacity-0 duration-300 p-2 select-none;
	}

	.clipboard-container:hover,
	.screen:hover + .clipboard-container {
		@apply opacity-[15%];
	}

	.clipboard-container > button {
		@apply border-none bg-inherit hover:cursor-pointer;
	}
</style>
