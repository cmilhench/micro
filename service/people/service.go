package people

import (
	"people/internal/uuid"
	"people/model"
	"people/repository"
)

type Service struct {
	Repo repository.Repository
}

func (s *Service) AddPerson(in *model.Person) (*model.Person, error) {
	id, err := uuid.New()
	if err != nil {
		return nil, err
	}
	in.ID = id
	out, err := s.Repo.Create(in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Service) GetPeople() ([]*model.Person, error) {
	return s.Repo.GetAll()
}

func (s *Service) GetPerson(id string) (*model.Person, error) {
	return s.Repo.GetByID(id)
}
