<script lang="ts">
	import { onMount } from "svelte";

  let data: {message?: string} = {};

  onMount(async() => {
    //let url = new URL("/api", window.location.href);
    // let test = await fetch(url)
    let test = await fetch("/api")
    data.message = await test.text()
  })
</script>

<svelte:head>
  <title>SPA</title>
</svelte:head>

<h1>Welcome to My App!</h1>
<p>Visit <a href="https://github.com/jackjohn7/go_svelte_spa_template">the project repo</a> to read the documentation</p>

{#if data.message}
  <p>{data.message}</p>
{/if}
