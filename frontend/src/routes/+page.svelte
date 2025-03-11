<script lang="ts">
	import {
		H1,
		P,
		ContentContainer,
		PageContainer,
		H2,
		AlbumCard,
		AlbumsContainer
	} from '$components';
	import AlbumImage from '$components/AlbumImage.svelte';

	const { data } = $props();
	const { albums } = data;
</script>

<PageContainer>
	<ContentContainer>
		<H1>Welcome to The Music Store!</H1>
		<P>Check out all of these great albums for sale (not really).</P>
	</ContentContainer>
	{#await albums}
		<P>Loading...</P>
	{:then albums}
		<AlbumsContainer>
			{#each albums as album (album.id)}
				<AlbumCard {album}>
					<H1>{album.title}</H1>
					<H2>{album.artist}</H2>
					<AlbumImage imageUrl={album.imageUrl} title={album.title} />
					<P>${album.price.toFixed(2)}</P>
				</AlbumCard>
			{/each}
		</AlbumsContainer>
	{:catch error}
		<P>{error.message}</P>
	{/await}
</PageContainer>
