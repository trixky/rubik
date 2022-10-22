<!-- ========================= HTML -->
<footer>
	<p>
		G0 -> G1
		<br/>
		orient edges
		<br/>
		<br/>
		G1 -> G2
		<br/>
		orient corners
		<br/>
		put FUDB edges in their slice
		<br/>
		<br/>
		G2 -> G3
		<br/>
		place corners
		<br/>
		put other edges in their slice
		<br/>
		<br/>
		G3 -> G4
		<br/>
		place edges
		<br/>
		<br/>
		Each group keeps the previous group rules
		<br/>
		by reducing the allowed moves.
	</p>
</footer>

<!-- ========================= CSS -->
<style lang="postcss">
	footer {
		@apply w-[350px] m-auto text-center pt-2 pb-10;
	}
</style>
