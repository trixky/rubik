<!-- ========================= SCRIPT -->
<script lang="ts">
	import './app.css';
	import Header from '../components/layout/header.svelte';
	import Footer from '../components/layout/footer.svelte';
</script>

<!-- ========================= HTML -->
<div id="chat" class="body-container">
	<Header />
	<main class="flex h-full">
		<slot />
	</main>
	<Footer />
</div>

<!-- ========================= CSS -->
<style lang="postcss">
	.body-container {
		@apply h-full;
	}
</style>
