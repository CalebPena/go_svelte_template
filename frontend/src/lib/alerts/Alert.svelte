<script lang="ts">
	import { fly } from 'svelte/transition';
	import type { Alert } from './alerts.svelte';
	import X from '$lib/icons/X.svelte';

	const { alert }: { alert: Alert } = $props();

	let message = $derived.by(() => {
		let formattedAmount: string;
		if (alert.count <= 1) {
			formattedAmount = '';
		} else if (alert.count <= 10) {
			formattedAmount = ' (' + alert.count + ')';
		} else {
			formattedAmount = ' (10+)';
		}

		return alert.message + formattedAmount;
	});

	let color = $derived.by(() => {
		if (alert.type === 'error') {
			return 'var(--primary-error)';
		}

		return 'var(--primary-color)';
	});
</script>

<div
	transition:fly={{ opacity: 1, x: '120%' }}
	role="alert"
	class="container"
	style:--color={color}
>
	<button
		onfocusin={() => {
			alert.pause();
		}}
		onfocusout={() => {
			alert.unpause();
		}}
		onclick={() => {
			alert.remove();
		}}
	>
		<X />
	</button>
	{message}
</div>

<style>
	.container {
		position: relative;
		width: 40em;
		max-width: 80vw;
		padding: var(--size-3);
		background-color: var(--primary-light);
		border: var(--border-size-3) solid var(--color);
		border-radius: var(--radius-3);
		word-break: break-word;
	}

	button {
		background: none;
		position: absolute;
		top: 0;
		right: 0;

		&:hover,
		&:focus-visible {
			background-color: var(--gray-3);
			border-radius: var(--radius-round);
		}
	}
</style>
