<script lang="ts">
	const { children } = $props();
</script>

<h2 class="text-center text-lg text-zinc-100">
	{@render children?.()}
</h2>
