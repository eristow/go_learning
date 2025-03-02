<script lang="ts">
	import { NavButton, P } from '$components';
	import H1 from '$components/H1.svelte';
	import '../app.css';
	let { children } = $props();
</script>

<div class="bg-zinc-800 p-4">
	<div class="grid grid-cols-3 items-center gap-4">
		<div>
			<NavButton href="/">Home</NavButton>
		</div>
		<H1>The Music Store</H1>
		<div></div>
	</div>
</div>
{@render children()}
