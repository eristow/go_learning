<script lang="ts">
	const { children } = $props();
</script>

<div
	class="m-auto rounded-md border-2 border-transparent bg-zinc-700 p-4 shadow-sm shadow-zinc-600"
>
	{@render children?.()}
</div>
