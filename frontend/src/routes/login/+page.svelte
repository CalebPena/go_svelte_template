<script lang="ts">
	import { alerts } from '$lib/alerts/alerts.svelte';
	import LoginForm from '$lib/auth/LoginForm.svelte';
	import type { PageProps } from './$types';
	import { enhance } from '$app/forms';
	import { loginPasswordSchema } from '$lib/schemas';
	import { handleActionWrapper } from '$lib/utils/forms';

	let { form }: PageProps = $props();

	$effect(() => {
		alerts.addAlertFromFormError(form);
	});

	let submitted = $state(false);
	let valid = $state(false);
</script>

<form
	method="POST"
	use:enhance={({ controller }) => {
		submitted = true;

		if (!valid) {
			controller.abort();
		}

		return handleActionWrapper();
	}}
>
	<LoginForm title="Login" {submitted} bind:valid passwordSchema={loginPasswordSchema}>
		{#snippet link()}<a href="/register" class="link">Click here to create an account</a>{/snippet}
	</LoginForm>
</form>
