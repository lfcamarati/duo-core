package entity

import "github.com/lfcamarati/duo-core/infra/database"

func NewService(title string, description *string, price float64) Service {
	return Service{
		Title:       title,
		Description: description,
		Price:       price,
	}
}

type Service struct {
	ID          *int64
	Title       string
	Description *string
	Price       float64
}

type ServiceRepository interface {
	database.Transactor

	Save(service Service) (*int64, error)
	Update(service Service) error
	GetAll() ([]Service, error)
	GetById(id int64) (*Service, error)
	Delete(id int64) error
}
