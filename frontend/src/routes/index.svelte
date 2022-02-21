<script lang="ts">
	import { onMount } from 'svelte';
	import type { EntityID } from '$lib/resources';
	import Entity from '$lib/components/Entity.svelte';
	import Group from '$lib/components/Group.svelte';
	import { browser } from '$app/env';

	export var entities = [];
	export var groups = [];
	let s: WebSocket;
	let wsConnected = false;
	let baseURL;

	if (browser) {
		baseURL = `${window.location.protocol}//${window.location.host}`;
	}

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
	onMount(() => {
		s = new WebSocket(`ws://${window.location.host}/ws`);
		console.log('WS connected', s);
		wsConnected = true;
		s.onmessage = (e: MessageEvent<WebSocketEventMap>) => {
			console.log(e);
			console.log(e.data);
			// TODO be smarter about what we fetch
			fetchData();
		};
		s.onclose = () => {
			console.log('WS Closed');
			wsConnected = false;
		};
	});

	function removeEntity(id: EntityID) {
		console.debug('remove entity: ', id);
		fetch(`${baseURL}/api/entity/${id}`, {
			method: 'DELETE'
		});
	}
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

<div class="mr-60 pr-60 w-full shadow-md bg-white dark:bg-slate-200 px-1">
	<h1 class="text-3xl font-bold">TF Sim</h1>
</div>

<div>
	<h2 class="text-2xl font-bold">Entities</h2>
	{#each entities as entity}
		<Entity {entity} remove={removeEntity} />
	{/each}
</div>

<div class="pb-10">
	<h2 class="text-2xl font-bold">Groups</h2>
	{#each groups as group}
		<Group {group} />
	{/each}
</div>

<div class="fixed bottom-0 p-2 grid grid-cols-3 text-sm font-medium bg-cyan-500 w-full">
	<div class="" />

	<div class="text-center font-bold">⚠ ALPHA ⚠</div>

	<div class="text-right ">
		{#if wsConnected}
			WS ⚡
		{/if}
	</div>
</div>
