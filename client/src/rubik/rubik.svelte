<!-- ========================= SCRIPT -->
<script lang="ts">
	import { onMount } from 'svelte';

	export let show = false;

	let canvas: any;

	onMount(async () => {
		const { createScene } = await import('./rubik');

		createScene(canvas);
	});
</script>

<!-- ========================= HTML -->
<canvas bind:this={canvas} class="rubik" style="opacity: {show ? '1' : '0'};" />

<!-- ========================= CSS -->
<style lang="postcss">
	.rubik {
		@apply absolute w-[266px] h-[266px] rounded-md duration-300;
	}
</style>
