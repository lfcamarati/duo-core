package usecase

import (
	"errors"

	"github.com/lfcamarati/duo-core/domain/clientpf/entity"
)

func NewDeleteClientPfUseCase(repository entity.ClientPfRepository) *DeleteClientPfUseCase {
	return &DeleteClientPfUseCase{repository}
}

type DeleteClientPfInput struct {
	ID int64
}

type DeleteClientPfOutput struct{}

type DeleteClientPfUseCase struct {
	Repository entity.ClientPfRepository
}

func (uc *DeleteClientPfUseCase) Execute(input DeleteClientPfInput) (*DeleteClientPfOutput, error) {
	clientPf, err := uc.Repository.GetById(input.ID)

	if err != nil {
		return nil, errors.New("cliente não encontrado")
	}

	if clientPf == nil {
		return nil, errors.New("cliente não encontrado")
	}

	err = uc.Repository.Delete(input.ID)

	if err != nil {
		return nil, err
	}

	return &DeleteClientPfOutput{}, nil
}
