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

	$: if (output_mode) {
		input_rubik?.visible(false);
		output_rubik?.visible(true);
	} else {
		input_rubik?.visible(true);
		output_rubik?.visible(false);
	}

	onMount(async () => {
		const engine = new Engine(266, 250, canvas);

		output_rubik = new Rubik(await generateCubes(), engine);
		output_rubik.visible(false);
		input_rubik = new Rubik(await generateCubes(), engine);
		output_rubik.visible(true);
	});
</script>

<!-- ========================= HTML -->
<canvas bind:this={canvas} class="rubik" style="opacity: {show ? '1' : '0'};" />

<!-- ========================= CSS -->
<style lang="postcss">
	.rubik {
		@apply absolute top-3 left-0 w-[266px] h-[266px] rounded-md duration-200;
	}
</style>
