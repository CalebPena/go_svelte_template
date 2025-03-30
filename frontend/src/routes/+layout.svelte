<script lang="ts">
	import Alerts from '$lib/alerts/Alerts.svelte';
	import './global.css';

	let { children } = $props();
</script>

<Alerts />

{@render children()}
