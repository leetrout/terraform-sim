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

<h1>TF Sim</h1>

<h1>Entities</h1>
<ul>
	{#each entities as entity}
		<li>{entity.Name}</li>
	{/each}
</ul>

<h1>Groups</h1>
<ul>
	{#each groups as group}
		<li>{group.Name}</li>
	{/each}
</ul>
