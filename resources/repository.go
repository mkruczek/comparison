package resources

import (
	"github.com/google/uuid"
)

type Repository interface {
	Create(r Resources) error
	Get(id uuid.UUID) (Resources, error)
	GetAll() ([]Resources, error)
	Update(r Resources) error
	Delete(id uuid.UUID) error
}

type inMemoryRepository struct {
	resources map[uuid.UUID]Resources
}

func NewRepository() Repository {
	return inMemoryRepository{
		resources: make(map[uuid.UUID]Resources),
	}
}

func (r inMemoryRepository) Create(res Resources) error {
	r.resources[res.ID] = res
	return nil
}

func (r inMemoryRepository) Get(id uuid.UUID) (Resources, error) {
	return r.resources[id], nil
}

func (r inMemoryRepository) GetAll() ([]Resources, error) {
	var resources []Resources
	for _, res := range r.resources {
		resources = append(resources, res)
	}
	return resources, nil
}

func (r inMemoryRepository) Update(res Resources) error {
	r.resources[res.ID] = res
	return nil
}

func (r inMemoryRepository) Delete(id uuid.UUID) error {
	delete(r.resources, id)
	return nil
}
