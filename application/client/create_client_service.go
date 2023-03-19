package client

import (
	"errors"

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

func (i *CreateClientCommand) Validate() error {
	if i.Name == nil || *i.Name == "" {
		return errors.New("nome deve ser informado")
	}

	if i.Type == nil || *i.Type == "" {
		return errors.New("tipo do cliente deve ser informado")
	}

	if i.CpfCnpj == nil || *i.CpfCnpj == "" {
		if client.ClientType(*i.Type) == client.PF {
			return errors.New("cpf deve ser informado")
		} else {
			return errors.New("cnpj deve ser informado")
		}
	}

	if i.Address == nil || *i.Address == "" {
		return errors.New("endereço deve ser informado")
	}

	if i.Email == nil || *i.Email == "" {
		return errors.New("email deve ser informado")
	}

	if i.Phone == nil || *i.Phone == "" {
		return errors.New("telefone deve ser informado")
	}

	return nil
}

type CreateClientService struct {
	NewRepository clientInfra.ClientRepositoryFactory
}

func (s *CreateClientService) Execute(input *CreateClientCommand) (*int64, error) {
	err := input.Validate()

	if err != nil {
		return nil, err
	}

	repository := s.NewRepository()

	if err := repository.Begin(); err != nil {
		return nil, err
	}
	defer repository.Rollback()

	clientPf := client.NewClient(*input.Name, *input.CpfCnpj, *input.Address, *input.Email, *input.Phone, client.ClientType(*input.Type))
	id, err := repository.Save(clientPf)

	if err != nil {
		return nil, err
	}

	if err := repository.Commit(); err != nil {
		return nil, err
	}

	return id, nil
}
