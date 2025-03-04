<script lang="ts">
	const { children, forField } = $props();
</script>

<label for={forField} class="text-zinc-100">
	{@render children?.()}
</label>
