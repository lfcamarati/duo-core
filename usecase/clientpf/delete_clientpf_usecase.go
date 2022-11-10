package usecase

import (
	"errors"

	"github.com/lfcamarati/duo-core/domain/clientpf/infra/repository"
)

func NewDeleteClientPfUseCase(factory repository.ClientPfRepositoryFactory) *DeleteClientPfUseCase {
	return &DeleteClientPfUseCase{factory}
}

type DeleteClientPfInput struct {
	ID int64
}

type DeleteClientPfOutput struct{}

type DeleteClientPfUseCase struct {
	NewRepository repository.ClientPfRepositoryFactory
}

func (uc *DeleteClientPfUseCase) Execute(input DeleteClientPfInput) (*DeleteClientPfOutput, error) {
	repository := uc.NewRepository()
	clientPf, err := repository.GetById(input.ID)

	if err != nil {
		return nil, errors.New("cliente não encontrado")
	}

	if clientPf == nil {
		return nil, errors.New("cliente não encontrado")
	}

	repository.Begin()
	err = repository.Delete(input.ID)

	if err != nil {
		repository.Rollback()
		return nil, err
	}

	repository.Commit()
	return &DeleteClientPfOutput{}, nil
}
