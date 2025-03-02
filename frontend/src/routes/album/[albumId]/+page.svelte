<script lang="ts">
	import { page } from '$app/state';
	import { ContentContainer, P, PageContainer, H1, H2, AlbumImage, NavButton } from '$components';
	const albumId = $derived(page.params.albumId);

	const { data } = $props();
	const { album } = data;
</script>

<PageContainer>
	<ContentContainer>
		{#await album}
			<P>Loading...</P>
		{:then album}
			<H1>{album.title}</H1>
			<H2>{album.artist}</H2>
			<P>${album.price.toFixed(2)}</P>
			<AlbumImage imageUrl={album.imageUrl} title={album.title} />
			<div class="grid grid-cols-2 gap-4">
				<NavButton href={`/album/${albumId}/edit`}>Edit</NavButton>
				<!-- TODO: change from NavButton to button with onClick to issue a delete -->
				<NavButton href={`/album/${albumId}/delete`}>Delete</NavButton>
			</div>
		{:catch error}
			<P>{error.message}</P>
		{/await}
	</ContentContainer>
</PageContainer>
