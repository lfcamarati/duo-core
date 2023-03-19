package client

import "github.com/lfcamarati/duo-core/infra/database"

type ClientRepository interface {
	database.Transactor

	GetAll() ([]Client, error)
	GetById(id int64) (*Client, error)
	Save(client Client) (*int64, error)
	Update(client Client) error
	Delete(id int64) error
}
