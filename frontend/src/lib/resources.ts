export type UUID = string;
export type EntityID = UUID;
export type GroupID = UUID;

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
