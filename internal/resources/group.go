package resources

import "github.com/google/uuid"

// Group allows grouping entities together in
// an EntitySet.
type Group struct {
	ID        string   // uuid
	Name      string   `validate:"required"`
	EntitySet []string `validate:"required,dive,uuid"`
}

func NewGroup(name string, entities []string) *Group {
	return &Group{
		ID:        uuid.NewString(),
		Name:      name,
		EntitySet: entities,
	}
}

type APICreateGroup struct {
	Name      string   `validate:"required"`
	EntitySet []string `validate:"required,dive,uuid"`
}

func (ag *APICreateGroup) AsGroup() *Group {
	return NewGroup(ag.Name, ag.EntitySet)
}
