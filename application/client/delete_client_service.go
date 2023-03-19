package client

import (
	"errors"

	clientInfra "github.com/lfcamarati/duo-core/infra/domain/client"
)

func NewDeleteClientService(factory clientInfra.ClientRepositoryFactory) *DeleteClientService {
	return &DeleteClientService{factory}
}

type DeleteClientCommand struct {
	ID int64
}

type DeleteClientService struct {
	NewRepository clientInfra.ClientRepositoryFactory
}

func (s *DeleteClientService) Execute(input DeleteClientCommand) error {
	repository := s.NewRepository()

	if err := repository.Begin(); err != nil {
		return err
	}
	defer repository.Rollback()

	clientPf, err := repository.GetById(input.ID)

	if err != nil {
		return errors.New("cliente não encontrado")
	}

	if clientPf == nil {
		return errors.New("cliente não encontrado")
	}

	err = repository.Delete(input.ID)

	if err != nil {
		return err
	}

	if err := repository.Commit(); err != nil {
		return err
	}

	return nil
}
