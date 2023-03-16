package resources

import "github.com/google/uuid"

type Resources struct {
	ID    uuid.UUID
	Name  string
	Value string
}
