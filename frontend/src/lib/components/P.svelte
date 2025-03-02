<script lang="ts">
	const { children } = $props();
</script>

<p class="text-center text-zinc-100">
	{@render children?.()}
</p>
