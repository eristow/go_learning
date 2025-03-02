<script lang="ts">
	let { children } = $props();
</script>

<button
	class="rounded border-2 border-zinc-700 bg-zinc-600 p-2 shadow-sm shadow-zinc-500 transition hover:border-zinc-100 hover:bg-zinc-800 hover:shadow-md"
>
	{@render children?.()}
</button>
