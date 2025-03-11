<script lang="ts">
	import { page } from '$app/state';
	import { NavButton } from '$components';
	import H1 from '$components/H1.svelte';
	import '../app.css';
	let { children } = $props();
</script>

<div class="bg-zinc-800 p-4">
	<div class="grid grid-cols-3 items-center gap-4">
		<div class="justify-self-start">
			<NavButton href="/">Home</NavButton>
		</div>
		<H1>The Music Store</H1>
		<div class="justify-self-end">
			{#if page.url.pathname == '/'}
				<NavButton href="/album/create">Add Album</NavButton>
			{/if}
		</div>
	</div>
</div>
{#if children}
	{@render children()}
{/if}
