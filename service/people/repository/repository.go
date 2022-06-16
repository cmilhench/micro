package repository

import "people/model"

type Repository interface {
	Create(*model.Person) (*model.Person, error)
	GetAll() ([]*model.Person, error)
	GetByID(string) (*model.Person, error)
	Update(*model.Person) (*model.Person, error)
	Delete(id string) error
}
