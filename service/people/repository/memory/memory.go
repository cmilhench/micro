package memory

import (
	"sync"

	"people/model"
)

type memoryRepository struct {
	sync.RWMutex
	people map[string]*model.Person
}

var DefaultRepository = &memoryRepository{
	people: make(map[string]*model.Person),
}

func New() *memoryRepository {
	return &memoryRepository{
		people: make(map[string]*model.Person),
	}
}

func (repo *memoryRepository) Create(person *model.Person) (*model.Person, error) {
	repo.Lock()
	defer repo.Unlock()

	if _, ok := repo.people[person.ID]; ok {
		return nil, model.ErrPersonAlreadyExists
	}
	repo.people[person.ID] = person
	return person, nil
}

func (repo *memoryRepository) GetAll() ([]*model.Person, error) {
	var values = make([]*model.Person, 0)
	for _, value := range repo.people {
		values = append(values, value)
	}
	return values, nil
}

func (repo *memoryRepository) GetByID(id string) (*model.Person, error) {
	if value, ok := repo.people[id]; ok {
		return value, nil
	}

	return nil, model.ErrPersonNotFound
}

func (repo *memoryRepository) Update(person *model.Person) (*model.Person, error) {
	repo.Lock()
	defer repo.Unlock()

	_, ok := repo.people[person.ID]
	if !ok {
		return nil, model.ErrPersonNotFound
	}
	repo.people[person.ID] = person
	return person, nil
}

func (repo *memoryRepository) Delete(id string) error {
	repo.Lock()
	defer repo.Unlock()

	if _, ok := repo.people[id]; !ok {
		return nil
	}
	delete(repo.people, id)
	return nil
}
