<!-- ========================= SCRIPT -->
<script lang="ts">
	import './app.css';
	import Header from '../components/layout/header.svelte';
	import Footer from '../components/layout/footer.svelte';
</script>

<!-- ========================= HTML -->
<Header />
<main class="flex h-full">
	<slot />
</main>
<Footer />
