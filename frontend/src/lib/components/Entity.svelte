<script lang="ts">
	import type { Entity, EntityID } from '$lib/resources';
	export let entity: Entity;
	export let remove: (id: EntityID) => void;
</script>

<div class="rounded-sm bg-slate-200 m-2 p-3 border-dotted border-2 border-black">
	<div>
		{entity.Name}
		<button type="button" class="text-xs button" on:click={() => remove(entity.ID)}
			>❌ Remove</button
		>
	</div>
	<div class="text-sm flex">
		<div class="flex-1">{entity.TurboEncabulationRate}</div>
		<div class="flex-1">{entity.RefractionRate}</div>
	</div>
	<div class="text-xs">{entity.ID}</div>
</div>
