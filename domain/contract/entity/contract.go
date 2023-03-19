package entity

import (
	"time"

	"github.com/lfcamarati/duo-core/infra/database"
)

func NewContract() *Contract {
	return &Contract{}
}

type Contract struct {
	ID        *int64
	Client    int64
	StartDate time.Time
	EndDate   time.Time
	Status    string
	Services  []ContractedService
}

type ContractedService struct {
	ID          *int64
	Title       string
	Description string
	Price       float64
}

type ContractRepository interface {
	database.Transactor

	Save(contract Contract) (*int64, error)
	Update(contract Contract) error
	GetAll() ([]Contract, error)
	GetById(id int64) (*Contract, error)
	Delete(id int64) error
}
