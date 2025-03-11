<script lang="ts">
	import type { MouseEventHandler } from 'svelte/elements';

	let {
		className = '',
		children,
		onClick = (() => {}) as MouseEventHandler<HTMLButtonElement>,
		type
	} = $props();
</script>

<button
	class={`cursor-pointer rounded border-2 border-zinc-700 bg-zinc-600 p-2 text-center text-zinc-100 shadow-sm shadow-zinc-500 transition hover:border-zinc-100 hover:bg-zinc-800 hover:text-zinc-200 hover:shadow-md ${className}`}
	onclick={onClick}
	{type}
>
	{@render children?.()}
</button>
