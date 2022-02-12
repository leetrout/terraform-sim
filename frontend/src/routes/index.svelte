<script context="module" lang="ts">
	export async function load({ fetch }) {
		// TODO where to store hostname config?
		const baseURL = 'http://localhost:9321';
		const [entitiesReq, groupsReq] = await Promise.all([
			fetch(`${baseURL}/api/entities`),
			fetch(`${baseURL}/api/groups`)
		]);
		let entities = [];
		let groups = [];
		if (entitiesReq.ok) {
			const entityJSONData = await entitiesReq.json();
			entities = entityJSONData;
		}
		if (groupsReq.ok) {
			const jsonData = await groupsReq.json();
			groups = jsonData;
		}
		return { props: { entities, groups } };
	}
</script>

<script lang="ts">
	export var entities;
	export var groups;
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
