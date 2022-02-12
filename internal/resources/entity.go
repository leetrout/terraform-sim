package resources

import "github.com/google/uuid"

// Entity represents a fixed collection of information
// that is a resource being managed by terraform.
// It has properties that allow the simple simulation of
// drift and slightly more complex state management compared
// to only using something like the random provider.
type Entity struct {
	ID                    string `validate:"required,uuid"`
	Name                  string `validate:"required"`
	TurboEncabulationRate int    `validate:"required"`
	RefractionRate        *int   // Nullable
}

func NewEntity(name string, ter int) *Entity {
	return &Entity{
		ID:                    uuid.NewString(),
		Name:                  name,
		TurboEncabulationRate: ter,
		RefractionRate:        nil,
	}
}

// APICreateEntity represents the data needed to create an entity
// via the API.
// ID is excluded.
type APICreateEntity struct {
	Name                  string `validate:"required"`
	TurboEncabulationRate int    `validate:"required"`
	RefractionRate        *int   // Nullable
}

// AsEntity returns a new Entity from the given APIEntity
func (ae *APICreateEntity) AsEntity() *Entity {
	e := NewEntity(ae.Name, ae.TurboEncabulationRate)
	if ae.RefractionRate != nil {
		e.RefractionRate = ae.RefractionRate
	}

	return e
}
