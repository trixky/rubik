<!-- ========================= SCRIPT -->
<script lang="ts">
	export let show = false;
</script>

<!-- ========================= HTML -->
<canvas class="rubik" style="opacity: {show ? '1' : '0'};" />

<!-- ========================= CSS -->
<style lang="postcss">
	.rubik {
		@apply absolute bg-black w-[266px] h-[266px] rounded-md;
	}
</style>
