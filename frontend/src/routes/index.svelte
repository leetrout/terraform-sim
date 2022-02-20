<script lang="ts">
	import { onMount } from 'svelte';

	export var entities = [];
	export var groups = [];

	// TODO move to stores?
	async function fetchData() {
		const baseURL = `${window.location.protocol}//${window.location.host}`;
		const [entitiesReq, groupsReq] = await Promise.all([
			fetch(`${baseURL}/api/entities`),
			fetch(`${baseURL}/api/groups`)
		]);
		if (entitiesReq.ok) {
			const entityJSONData = await entitiesReq.json();
			entities = entityJSONData;
		}
		if (groupsReq.ok) {
			const groupJSONData = await groupsReq.json();
			groups = groupJSONData;
		}
	}

	onMount(fetchData);
</script>

<!-- TODO Add nav -->
<!-- <div class="w-60 h-full shadow-md bg-white px-1 absolute">
	<ul class="relative">
		<li class="relative">
			<a
				class="flex items-center text-sm py-4 px-6 h-12 overflow-hidden text-gray-700 text-ellipsis whitespace-nowrap rounded hover:text-gray-900 hover:bg-gray-100 transition duration-300 ease-in-out"
				href="#!"
				data-mdb-ripple="true"
				data-mdb-ripple-color="dark">Sidenav link 1</a
			>
		</li>
		<li class="relative">
			<a
				class="flex items-center text-sm py-4 px-6 h-12 overflow-hidden text-gray-700 text-ellipsis whitespace-nowrap rounded hover:text-gray-900 hover:bg-gray-100 transition duration-300 ease-in-out"
				href="#!"
				data-mdb-ripple="true"
				data-mdb-ripple-color="dark">Sidenav link 2</a
			>
		</li>
		<li class="relative">
			<a
				class="flex items-center text-sm py-4 px-6 h-12 overflow-hidden text-gray-700 text-ellipsis whitespace-nowrap rounded hover:text-gray-900 hover:bg-gray-100 transition duration-300 ease-in-out"
				href="#!"
				data-mdb-ripple="true"
				data-mdb-ripple-color="dark">Sidenav link 2</a
			>
		</li>
	</ul>
</div> -->

<div class="mr-60 pr-60 w-full shadow-md bg-white px-1">
	<h1 class="text-3xl font-bold">TF Sim</h1>

	<h2 class="text-2xl font-bold">Entities</h2>
	<ul class="pl-5 list-disc">
		{#each entities as entity}
			<li class="">{entity.Name}</li>
		{/each}
	</ul>

	<h2 class="text-2xl font-bold">Groups</h2>
	<ul class="pl-5 list-disc">
		{#each groups as group}
			<li>{group.Name}</li>
		{/each}
	</ul>
</div>
