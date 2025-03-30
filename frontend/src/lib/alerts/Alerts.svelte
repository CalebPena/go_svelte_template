<script lang="ts">
	import Alert from './Alert.svelte';
	import { alerts } from './alerts.svelte';
</script>

<!-- svelte-ignore a11y_mouse_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	onmouseover={() => {
		alerts.pauseAll();
	}}
	onmouseleave={() => {
		alerts.unpauseAll();
	}}
>
	{#each alerts.alerts as alert (alert.id)}
		<Alert {alert} />
	{/each}
</div>

<style>
	div {
		position: fixed;
		right: 0;
		display: flex;
		flex-direction: column;
		gap: var(--size-3);
		padding: var(--size-3);
	}
</style>
