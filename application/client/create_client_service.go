package client

import (
	"github.com/lfcamarati/duo-core/domain/client"
	clientInfra "github.com/lfcamarati/duo-core/infra/domain/client"
)

func NewCreateClientService(factory clientInfra.ClientRepositoryFactory) *CreateClientService {
	return &CreateClientService{factory}
}

type CreateClientCommand struct {
	Name    *string
	CpfCnpj *string
	Address *string
	Email   *string
	Phone   *string
	Type    *string
}

type CreateClientService struct {
	NewRepository clientInfra.ClientRepositoryFactory
}

func (s *CreateClientService) Execute(input *CreateClientCommand) (*int64, error) {
	repository := s.NewRepository()

	if err := repository.Begin(); err != nil {
		return nil, err
	}
	defer repository.Rollback()

	client, err := client.NewClient(input.Name, input.CpfCnpj, input.Address, input.Email, input.Phone, input.Type)

	if err != nil {
		return nil, err
	}

	id, err := repository.Save(client)

	if err != nil {
		return nil, err
	}

	if err := repository.Commit(); err != nil {
		return nil, err
	}

	return id, nil
}
