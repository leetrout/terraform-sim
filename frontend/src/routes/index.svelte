<script lang="ts">
	export var entities = [];
	export var groups = [];
	const baseURL = 'http://localhost:9321';

	// TODO move to stores?
	async function fetchData() {
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

	fetchData();
</script>

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
