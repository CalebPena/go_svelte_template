<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { alerts } from '$lib/alerts/alerts.svelte';
	import { DEFUALT_ERROR_MESSAGE, HOME_PAGE_PATH, REDIRECT_PARAM_NAME } from '$lib/constants';
	import { onMount } from 'svelte';

	onMount(() => {
		if (page.error !== null) {
			alerts.addAlert(page.error.message, 'error');

			if (page.status === 403) {
				const params = new URLSearchParams();
				const currentPath = window.location.pathname;

				if (currentPath !== HOME_PAGE_PATH && currentPath !== '') {
					params.set(REDIRECT_PARAM_NAME, currentPath);
					goto(`/login?${String(params)}`);
					return;
				}

				goto('/login');
			} else if (page.status === 404) {
				goto(HOME_PAGE_PATH);
			}
			return;
		}

		alerts.addAlert(DEFUALT_ERROR_MESSAGE, 'error');
	});
</script>

<h1>
	{#if page.error !== null}
		{page.error.message}
	{:else}
		{DEFUALT_ERROR_MESSAGE}
	{/if}
</h1>

<style>
	h1 {
		text-align: center;
		min-width: 100%;
		padding: var(--size-7);
	}
</style>
