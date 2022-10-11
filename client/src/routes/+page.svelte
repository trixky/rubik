<!-- ========================= SCRIPT -->
<script lang="ts">
	import ResultStore from '../stores/result';
	import ApiPostResolve from '../api/post.resolve';
	import * as StaticInstructions from '../stores/static/instructions';
	import SanitizeInput from '../sanitizers/input';

	const screen_rows = 8;
	const screen_columns = 8;
	const max_instructions = screen_rows * screen_columns;
	const max_selected_input = max_instructions - 1;

	$: output_mode = $ResultStore.length > 0;
	let rubik_mode = false;

	let prompt_id = 0;
	let prompt_period = false;
	let inputs: string[] = [];
	let input_selected = 0;
	let output_selected = 0;

	$: selected = output_mode ? output_selected : input_selected;

	$: end_selected = output_mode
		? output_selected === $ResultStore.length
		: input_selected === inputs.length;
	$: last_input_selected = output_mode
		? output_selected === $ResultStore.length - 1
		: input_selected === inputs.length - 1;
	$: no_inputs = output_mode ? $ResultStore.length === 0 : inputs.length === 0;
	$: input_str = inputs.join(' ');

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
			const input_selected_value = inputs[input_selected];
			if (input_selected_value != undefined && input_selected_value[0] === instruction[0]) {
				const base = instruction[0];
				switch ((input_selected_value + instruction).replaceAll(base, 'X')) {
					case 'XX':
					case "X'X'":
						inputs[input_selected] = base + '2';
						break;
					case 'X2X':
						inputs[input_selected] = base + "'";
						break;
					case "X2X'":
					case "X'X":
						inputs[input_selected] = base;
						break;
					default:
						inputs[input_selected] = instruction;
						break;
				}
			} else if (!end_selected) {
				inputs[input_selected] = instruction;
			} else if (inputs.length < max_instructions) {
				inputs = [...inputs, instruction];
			}
			handleHorizontalMove(true);
		}
	}

	function handleHorizontalMove(direction: boolean) {
		new_prompte();
		if (direction) {
			// right
			output_mode
				? (output_selected = selectSafe(output_selected + 1))
				: (input_selected = selectSafe(input_selected + 1));
		} else {
			// left
			if (output_mode) {
				output_selected = selectSafe(output_selected - 1);
			} else {
				input_selected = selectSafe(input_selected - 1);
			}
		}
	}

	function handleHorizontalSuperMove(direction: boolean) {
		new_prompte();
		if (direction) {
			// end
			if (output_mode) {
				output_selected = selectSafe($ResultStore.length);
			} else {
				input_selected = selectSafe(inputs.length);
			}
		} else {
			// start
			if (output_mode) {
				output_selected = selectSafe(0);
			} else {
				input_selected = selectSafe(0);
			}
		}
	}

	function handleVerticalMove(direction: boolean) {
		new_prompte();
		if (direction) {
			// up
			if (output_mode) {
				if (output_selected >= screen_rows)
					output_selected = selectSafe(output_selected - screen_rows);
			} else {
				if (input_selected >= screen_rows)
					input_selected = selectSafe(input_selected - screen_rows);
			}
		} else {
			// down
			if (output_mode) {
				if (output_selected < $ResultStore.length - screen_rows)
					output_selected = selectSafe(output_selected + screen_rows);
			} else {
				if (input_selected < inputs.length - screen_rows + 1)
					input_selected = selectSafe(input_selected + screen_rows);
			}
		}
	}

	function handleReset() {
		input_selected = selectSafe(0);
		output_selected = selectSafe(0);
		inputs = [];
		ResultStore.reset();
		new_prompte();
	}

	function handleInsert() {
		inputs = [
			...inputs.slice(0, input_selected),
			StaticInstructions.instructions[0],
			...inputs.slice(input_selected)
		];
	}

	function handleDelete() {
		if (output_mode) {
			output_selected = 0;
			ResultStore.reset();
		} else {
			if (last_input_selected) {
				inputs = inputs.slice(0, input_selected);
			} else if (!end_selected) {
				inputs = [...inputs.slice(0, input_selected + 1), ...inputs.slice(input_selected + 2)];
			} else if (!no_inputs) {
				inputs = inputs.slice(0, input_selected - 1);
				handleHorizontalMove(false);
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
		<div class="screen" style="opacity: {rubik_mode ? '0.1' : '1'};">
			{#each output_mode ? $ResultStore : inputs as instruction, index}
				<p class:selected-input={index === selected && prompt_period} class="screen-instruction">
					{instruction}
				</p>
			{/each}
			{#if !output_mode && inputs.length < max_instructions}
				<span class:selected-input={end_selected && prompt_period} class="input">&nbsp;</span>
			{/if}
		</div>
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
			<button
				class="physic-button left-rotation special-instruction text-red-400"
				on:click={handleReset}>rst</button
			>
			<button
				class="physic-button left-rotation special-instruction text-red-400"
				on:click={handleDelete}>del</button
			>
			<button class="physic-button left-rotation special-instruction" on:click={handleInsert}
				>ins</button
			>
			<button class="physic-button right-rotation special-instruction">grp</button>
			<button class="physic-button right-rotation special-instruction" on:click={handleRandom}
				>ran</button
			>
			<button class="physic-button right-rotation special-instruction">on</button>
			{#each StaticInstructions.physical_instructions as instruction, index}
				<button
					class="physic-button {index % 6 >= 3 ? 'right-rotation' : 'left-rotation'}"
					on:click={() => handleInstruction(instruction)}>{instruction.toLocaleLowerCase()}</button
				>
			{/each}
			<button
				class="physic-button left-rotation special-instruction"
				on:click={() => handleVerticalMove(false)}>{'v'}</button
			>
			<button
				class="physic-button left-rotation special-instruction"
				on:click={() => handleHorizontalSuperMove(false)}>{'<<'}</button
			>
			<button
				class="physic-button left-rotation special-instruction"
				on:click={() => handleHorizontalMove(false)}>{'<'}</button
			>
			<button
				class="physic-button right-rotation special-instruction"
				on:click={() => handleHorizontalMove(true)}>{'>'}</button
			>
			<button
				class="physic-button right-rotation special-instruction"
				on:click={() => handleHorizontalSuperMove(true)}>{'>>'}</button
			>
			<button
				class="physic-button right-rotation special-instruction"
				style="rotate: 180deg;"
				on:click={() => handleVerticalMove(true)}>{'v'}</button
			>
			<button class="physic-button left-rotation special-instruction" style="rotate: 180deg;"
				>...</button
			>
			<button class="physic-button left-rotation special-instruction" style="rotate: 180deg;"
				>...</button
			>
			<button class="physic-button left-rotation special-instruction" style="rotate: 180deg;"
				>...</button
			>
			<button class="physic-button right-rotation special-instruction" style="rotate: 180deg;"
				>...</button
			>
			<button class="physic-button right-rotation special-instruction" on:click={handleDimension}
				>dim</button
			>
			<button class="physic-button right-rotation special-instruction" on:click={handleResolve}
				>ok</button
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
		@apply grid grid-cols-8 grid-rows-8 m-0 w-[240px] h-[240px] p-3 border-solid border-[1px] border-black rounded-md break-words duration-200;
		box-shadow: inset 0 0 10px rgba(0, 0, 0, 0.322);
		background-color: #c9e9c5;
	}

	.physic-button-container {
		@apply grid grid-cols-6 grid-rows-5 gap-2 mt-5 mb-0;
	}

	.physic-button-container > button {
		@apply border-solid border-[1px] border-black rounded-md;
	}

	.right-rotation {
		transform: rotate(8deg);
	}

	.left-rotation {
		transform: rotate(-8deg);
	}

	.physic-button {
		@apply border-none px-[7px] py-[8px] hover:cursor-pointer;
		font-weight: 500;
		background-color: azure;
	}

	.physic-button:hover {
		filter: brightness(92%);
	}

	.physic-button:hover:active {
		filter: brightness(80%);
	}

	.special-instruction {
		background-color: #eee;
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
