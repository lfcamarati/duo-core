package service

import "github.com/lfcamarati/duo-core/infra/database"

type ServiceRepository interface {
	database.Transactor

	Save(service Service) (*int64, error)
	Update(service Service) error
	GetAll() ([]Service, error)
	GetById(id int64) (*Service, error)
	Delete(id int64) error
}
