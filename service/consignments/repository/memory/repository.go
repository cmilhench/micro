package memory

import (
	"strconv"
	"sync"

	"consignments/model"
)

type memoryRepository struct {
	sync.RWMutex
	consignments []*model.Consignment
}

func New() *memoryRepository {
	return &memoryRepository{
		consignments: make([]*model.Consignment, 0, 0),
	}
}

func (repo *memoryRepository) Create(consignment *model.Consignment) (*model.Consignment, error) {
	repo.Lock()
	defer repo.Unlock()
	consignment.ID = strconv.Itoa(len(repo.consignments) + 1)
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *memoryRepository) GetAll() []*model.Consignment {
	return repo.consignments
}
