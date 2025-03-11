<script lang="ts">
	type InputProps = {
		type: string;
		id: string;
		name: string;
		required: boolean;
		value?: string | number;
		step?: string;
	};

	const { type, id, name, value = '', required, step = '0' }: InputProps = $props();
</script>

<input class="rounded-lg" {type} {id} {name} {value} {required} {step} />
