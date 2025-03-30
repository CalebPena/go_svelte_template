<script lang="ts">
	import { formInputError } from '$lib/utils/forms';
	import { emailSchema } from '$lib/schemas';
	import type { Snippet } from 'svelte';
	import { Schema } from 'zod';
	let {
		title,
		link,
		submitted,
		passwordSchema,
		valid = $bindable(),
	}: {
		title: string;
		link: Snippet;
		submitted: boolean;
		passwordSchema: Schema;
		valid: boolean;
	} = $props();

	let email = $state('');
	let password = $state('');

	let emailErrorMessage = $derived(formInputError(emailSchema, email));
	let passwordErrorMessage = $derived(formInputError(passwordSchema, password));

	$effect(() => {
		valid = emailErrorMessage === '' && passwordErrorMessage === '';
	});
</script>

<div class="center">
	<div class="container">
		<h1>{title}</h1>
		<div class="input-container">
			<div>
				<label for="email">Email:</label>
			</div>
			<div>
				<input type="email" id="email" name="email" bind:value={email} />
			</div>
			{#if emailErrorMessage !== '' && submitted}
				<p class="err-msg">{emailErrorMessage}</p>
			{/if}
			<div>
				<label for="password">Password:</label>
			</div>
			<div>
				<input type="password" id="password" name="password" bind:value={password} />
			</div>
			{#if passwordErrorMessage !== '' && submitted}
				<p class="err-msg">{passwordErrorMessage}</p>
			{/if}
			<div>
				{@render link()}
			</div>
			<div class="button-container">
				<button type="submit" class="btn">Login</button>
			</div>
		</div>
	</div>
</div>

<style>
	.center {
		display: flex;
		align-items: center;
		justify-content: center;
		flex-direction: column;
		height: 80vh;
	}

	.container {
		background-color: var(--primary-background-color);
		border: var(--border-size-3) solid var(--primary-color);
		border-radius: var(--radius-3);
		box-shadow: var(--shadow-4);
	}

	h1 {
		font-size: 2.5em;
		font-weight: var(--font-weight-6);
		margin: 0;
		text-align: center;
		background-color: var(--primary-color);
		color: var(--primary-light);
		padding: var(--size-relative-1) 0 var(--size-relative-2) 0;
		min-width: 100%;
	}

	.input-container {
		padding: var(--size-relative-5) var(--size-relative-10) var(--size-relative-10)
			var(--size-relative-10);
	}

	.button-container {
		padding: var(--size-5) 0;
	}

	@media (max-width: 450px) {
		.center {
			font-size: var(--font-size-0);
		}
	}
</style>
