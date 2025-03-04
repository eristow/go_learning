<script lang="ts">
	import { page } from '$app/state';
	import {
		ContentContainer,
		P,
		PageContainer,
		H1,
		H2,
		AlbumImage,
		NavButton,
		Button
	} from '$components';
	const albumId = $derived(page.params.albumId);

	const { data } = $props();
	// const { album, deleteAlbum } = data;
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
				<form class="w-full" method="POST" action={`/album/${albumId}?/delete`}>
					<input type="hidden" name="_method" value="DELETE" />
					<Button className="w-full" type="submit">Delete</Button>
				</form>
			</div>
		{:catch error}
			<P>{error.message}</P>
		{/await}
	</ContentContainer>
</PageContainer>
