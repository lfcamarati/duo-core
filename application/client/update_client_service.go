package client

import (
	"errors"

	"github.com/lfcamarati/duo-core/domain/client"
	clientInfra "github.com/lfcamarati/duo-core/infra/domain/client"
)

func NewUpdateClientService(factory clientInfra.ClientRepositoryFactory) *UpdateClientService {
	return &UpdateClientService{factory}
}

type UpdateClientCommand struct {
	ID      int64
	Name    *string
	CpfCnpj *string
	Address *string
	Email   *string
	Phone   *string
	Type    *string
}

func (i *UpdateClientCommand) Validate() error {
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

type UpdateClientService struct {
	NewRepository clientInfra.ClientRepositoryFactory
}

func (s *UpdateClientService) Execute(command *UpdateClientCommand) error {
	repository := s.NewRepository()

	if err := repository.Begin(); err != nil {
		return err
	}
	defer repository.Rollback()

	savedClient, err := repository.GetById(command.ID)

	if err != nil {
		return err
	}

	savedClient.Name = *command.Name
	savedClient.CpfCnpj = *command.CpfCnpj
	savedClient.Address = *command.Address
	savedClient.Email = *command.Email
	savedClient.Phone = *command.Phone
	savedClient.Type = client.ClientType(*command.Type)

	err = repository.Update(*savedClient)

	if err != nil {
		return err
	}

	if err := repository.Commit(); err != nil {
		return err
	}

	return nil
}
