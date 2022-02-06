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

func NewEntity() *Entity {
	return &Entity{
		ID:                    uuid.NewString(),
		Name:                  "",
		TurboEncabulationRate: 0,
		RefractionRate:        nil,
	}
}
