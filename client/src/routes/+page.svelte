<!-- ========================= SCRIPT -->
<script lang="ts">
    import ResultStore from '../stores/result'
    import ApiPostResolve from '../api/post.resolve'

    let input = ""

    function handleResolve() {
        ApiPostResolve(input).then(result => {
            ResultStore.set(result)
        })
    }
</script>
<!-- ========================= HTML -->
<h1>Rubik</h1>
<div>
    <h2>Input</h2>
    <input class="output" type="text" bind:value={input} >
    <button on:click={handleResolve}>resolve</button>
</div>
<div>
    <h2>Output</h2>
    <p class="output">{$ResultStore}</p>
</div>

<!-- ========================= CSS -->
<style lang="postcss">
    .output {
        display: inline-block;
        padding: 10px;
        min-width: 300px;
        min-height: 100px;
        border: black solid 1px;
    }
</style>


