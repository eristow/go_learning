<script lang="ts">
	import { Button, Input, Label, P } from '$components';
	import ContentContainer from '$components/ContentContainer.svelte';
	import H1 from '$components/H1.svelte';
	import PageContainer from '$components/PageContainer.svelte';

	const { data } = $props();
	const { album } = data;
</script>

<PageContainer>
	<ContentContainer>
		{#await album}
			<P>Loading...</P>
		{:then album}
			<form action={`/album/${album.id}?/edit`} method="POST">
				<H1>Album Edit</H1>
				<div class="flex flex-col space-y-4">
					<div>
						<Label forField="albumTitle">Album Title:</Label>
						<Input type="text" id="albumTitle" name="albumTitle" value={album.title} required />
					</div>

					<div>
						<Label forField="albumArtist">Artist:</Label>
						<Input type="text" id="albumArtist" name="albumArtist" value={album.artist} required />
					</div>

					<div>
						<Label forField="albumPrice">Price:</Label>
						<Input
							type="number"
							step="0.01"
							id="albumPrice"
							name="albumPrice"
							value={album.price}
							required
						/>
					</div>

					<div>
						<Label forField="albumArt">Album Art URL:</Label>
						<!-- TODO: load current album art, get new album art, then upload to S3 -->
						<!-- <Input type="file" id="albumArt" name="albumArt" accept="image/*" required /> -->
						<Input type="text" id="albumArt" name="albumArt" value={album.imageUrl} required />
					</div>

					<Button type="submit">Edit Album</Button>
				</div>
			</form>
		{:catch error}
			<P>{error.message}</P>
		{/await}
	</ContentContainer>
</PageContainer>
