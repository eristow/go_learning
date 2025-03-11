<script>
	let { children, href } = $props();
</script>

<a
	{href}
	class="rounded border-2 border-zinc-700 bg-zinc-600 p-2 text-center text-zinc-100 shadow-sm shadow-zinc-500 transition hover:border-zinc-100 hover:bg-zinc-800 hover:text-zinc-200 hover:shadow-md"
>
	{@render children?.()}
</a>
