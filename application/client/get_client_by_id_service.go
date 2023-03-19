package client

import (
	"database/sql"
	"errors"

	"github.com/lfcamarati/duo-core/domain/client"
	clientInfra "github.com/lfcamarati/duo-core/infra/domain/client"
)

var (
	ErrClientNotFound = errors.New("cliente nao encontrado")
)

func NewGetClientByIdService(factory clientInfra.ClientRepositoryFactory) GetClientByIdService {
	return GetClientByIdService{factory}
}

type GetClientByIdService struct {
	NewRepository clientInfra.ClientRepositoryFactory
}

func (service *GetClientByIdService) Execute(id int64) (*client.Client, error) {
	repository := service.NewRepository()
	client, err := repository.GetById(id)

	if err == sql.ErrNoRows {
		return nil, ErrClientNotFound
	}

	return client, nil
}
