import * as yup from 'yup';

export type UUID = string;
export type EntityID = UUID;
export type GroupID = UUID;

// TODO: consider moving to Zod?
// https://github.com/colinhacks/zod

export const EntitySchema = yup
	.object({
		Name: yup.string().required(),
		TurboEncabulationRate: yup.number().required(),
		RefractionRate: yup.number().label('Refraction Rate')
	})
	.required();

export type CreateEntity = yup.InferType<typeof EntitySchema>;

// TODO: Can you combine the inferred type? only need to add the entity ID
export interface Entity {
	ID: EntityID;
	Name: string;
	TurboEncabulationRate: number;
	RefractionRate?: number;
}

export type EntityAPI = Omit<Entity, 'ID'>;

export interface Group {
	ID: GroupID;
	Name: string;
	EntitySet: EntityID[];
}

export type GroupAPI = Omit<Group, 'ID'>;
