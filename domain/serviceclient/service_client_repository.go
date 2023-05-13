package serviceclient

import "github.com/lfcamarati/duo-core/infra/database"

type ServiceClientRepository interface {
	database.Transactor

	Register(serviceClient ServiceClient) (*int64, error)
}
