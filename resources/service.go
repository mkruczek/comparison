package resources

import "github.com/google/uuid"

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) Create(r Resources) (Resources, error) {
	r.ID = uuid.New()
	if err := s.repo.Create(r); err != nil {
		return Resources{}, err
	}

	return r, nil
}

func (s Service) Get(id uuid.UUID) (Resources, error) {
	return s.repo.Get(id)
}

func (s Service) GetAll() ([]Resources, error) {
	return s.repo.GetAll()
}

func (s Service) Update(r Resources) (Resources, error) {
	return r, s.repo.Update(r)
}

func (s Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
