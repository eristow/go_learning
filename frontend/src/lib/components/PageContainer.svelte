<script lang="ts">
	const { children } = $props();
</script>

<div class="flex flex-col bg-zinc-900 p-4">
	{@render children?.()}
</div>
