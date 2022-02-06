package resources

import "github.com/google/uuid"

// Entity represents a fixed collection of information
// that is a resource being managed by terraform.
// It has properties that allow the simple simulation of
// drift and slightly more complex state management compared
// to only using something like the random provider.
type Entity struct {
	ID                    uuid.UUID
	Name                  string
	TurboEncabulationRate int
	RefractionRate        *int // Nullable
}

func NewEntity() *Entity {
	return &Entity{
		ID:                    uuid.New(),
		Name:                  "",
		TurboEncabulationRate: 0,
		RefractionRate:        nil,
	}
}
