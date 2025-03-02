<script>
	import { Button } from '$components';

	let { children, href } = $props();
</script>

<a
	{href}
	class="rounded border-2 border-neutral-700 bg-neutral-600 p-2 text-center text-neutral-100 shadow-sm shadow-neutral-500 transition hover:border-neutral-100 hover:bg-neutral-800 hover:text-neutral-200 hover:shadow-md"
>
	{@render children?.()}
</a>
