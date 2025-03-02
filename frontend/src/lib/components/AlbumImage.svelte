<script lang="ts">
	const { imageUrl, title } = $props();
</script>

<img
	src={imageUrl}
	alt={title}
	class="mx-auto my-2 h-auto max-h-[300px] w-full max-w-[300px] rounded-lg"
/>
