<script lang="ts">
	const { children } = $props();
</script>

<div class="grid grid-cols-3 gap-4 pt-4">
	{@render children?.()}
</div>
