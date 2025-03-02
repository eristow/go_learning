<script lang="ts">
	const { children } = $props();
</script>

<h1 class="text-center text-2xl font-bold text-zinc-100">
	{@render children?.()}
</h1>
