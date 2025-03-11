<script lang="ts">
	import { Button, Input, Label } from '$components';
	import ContentContainer from '$components/ContentContainer.svelte';
	import H1 from '$components/H1.svelte';
	import PageContainer from '$components/PageContainer.svelte';
</script>

<PageContainer>
	<ContentContainer>
		<H1>Album Create</H1>
		<form action="/album?/create" method="POST" data-testid="form">
			<div class="flex flex-col space-y-4">
				<div>
					<Label forField="albumTitle">Album Title:</Label>
					<Input type="text" id="albumTitle" name="albumTitle" required />
				</div>

				<div>
					<Label forField="albumArtist">Artist:</Label>
					<Input type="text" id="albumArtist" name="albumArtist" required />
				</div>

				<div>
					<Label forField="albumPrice">Price:</Label>
					<Input type="number" step="0.01" id="albumPrice" name="albumPrice" required />
				</div>

				<div>
					<Label forField="albumArt">Album Art URL:</Label>
					<!-- TODO: get new album art, then upload to S3 -->
					<!-- <Input type="file" id="albumArt" name="albumArt" accept="image/*" required /> -->
					<Input type="text" id="albumArt" name="albumArt" required />
				</div>

				<Button type="submit">Create Album</Button>
			</div>
		</form>
	</ContentContainer>
</PageContainer>
