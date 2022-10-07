<!-- ========================= SCRIPT -->
<script lang="ts">
	import './app.css';
	import Header from '../components/layout/header.svelte';
	import Footer from '../components/layout/footer.svelte';
</script>

<!-- ========================= HTML -->
<Header />
<main>
	<slot />
</main>
<Footer />
