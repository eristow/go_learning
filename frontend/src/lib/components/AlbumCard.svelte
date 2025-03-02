<script lang="ts">
	const { children, album } = $props();
</script>

<div
	class="rounded-lg border-2 border-transparent bg-zinc-700 p-4 shadow-sm shadow-zinc-600 transition-shadow hover:border-zinc-600 hover:shadow-lg"
>
	<a href={`/album/${album.id}`}>
		{@render children?.()}
	</a>
</div>
