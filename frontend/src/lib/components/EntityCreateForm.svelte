<script lang="ts">
	import { createForm } from 'felte';
	import { validator } from '@felte/validator-yup';
	import * as yup from 'yup';
	import { browser } from '$app/env';

	import { EntitySchema } from '$lib/resources';
	import type { CreateEntity } from '$lib/resources';

	let baseURL: string;

	// FIXME: don't duplicate
	if (browser) {
		baseURL = `${window.location.protocol}//${window.location.host}`;
	}

	const { form, errors, reset } = createForm({
		extend: validator<CreateEntity>({ schema: EntitySchema }),
		onSubmit: async (values) => {
			const payload = JSON.stringify(values);

			let resp = fetch(`${baseURL}/api/entity`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: payload
			});

			// TODO: do something with the response
			resp.then((x) => console.log(x));

			if ((await resp).status == 201) {
				reset();
			}
		}
	});
</script>

<div class="border mb-8 p-4">
	<form class="form" use:form on:submit|preventDefault>
		<div class="mb-4">
			<label class="block text-gray-700 text-sm font-bold mb-2" for="Name"> Entity Name </label>
			<input
				class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
				id="Name"
				type="text"
				name="Name"
				placeholder="Entity Name"
			/>
			{#if $errors.Name}
				<span class="error">{$errors.Name}</span>
			{/if}
		</div>

		<div class="mb-4">
			<label class="block text-gray-700 text-sm font-bold mb-2" for="TurboEncabulationRate">
				Turbo Encabulation Rate
			</label>
			<input
				class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
				id="TurboEncabulationRate"
				type="number"
				name="TurboEncabulationRate"
				placeholder="1"
			/>
			{#if $errors.TurboEncabulationRate}
				<span class="error">{$errors.TurboEncabulationRate}</span>
			{/if}
		</div>
		<div class="mb-4">
			<label class="block text-gray-700 text-sm font-bold mb-2" for="RefractionRate">
				Refraction Rate
			</label>
			<input
				class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
				id="RefractionRate"
				type="number"
				name="RefractionRate"
				placeholder="Optional"
			/>
			{#if $errors.RefractionRate}
				<span class="error">{$errors.RefractionRate}</span>
			{/if}
		</div>
		<button
			class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
			type="submit"
		>
			Add Entity
		</button>
	</form>
</div>
