<!-- ========================= SCRIPT -->
<script lang="ts">
	import { onMount } from 'svelte';
	import { Rubik } from './rubik';
	import Engine from './engine';
	import generateCubes from './cubes';

	export let output_mode = false;
	export let show = false;
	export let input_rubik: Rubik | undefined = undefined;
	export let output_rubik: Rubik | undefined = undefined;

	let canvas: any;

	let engine: Engine;

	$: engine?.setCameraMovement(show);

	$: if (output_mode) {
		// if output mode
		// hide the input rubik
		input_rubik?.visible(false);
		// show the output rubik
		output_rubik?.visible(true);
	} else {
		// if input mode
		// show the input rubik
		input_rubik?.visible(true);
		// hide the output rubik
		output_rubik?.visible(false);
	}

	onMount(async () => {
		engine = new Engine(266, 250, canvas);

		// create the output rubik
		output_rubik = new Rubik(await generateCubes(), engine);
		// hide the  the output rubik
		output_rubik.visible(false);
		// create the input rubik
		input_rubik = new Rubik(await generateCubes(), engine);
		// show the  the input rubik
		input_rubik.visible(false);
	});
</script>

<!-- ========================= HTML -->
<canvas bind:this={canvas} class="rubik" class:show />

<!-- ========================= CSS -->
<style lang="postcss">
	.rubik {
		@apply absolute top-3 left-0 w-[266px] h-[266px] rounded-md duration-300 opacity-0;
	}

	.rubik.show {
		@apply hover:cursor-pointer opacity-100;
	}
</style>
