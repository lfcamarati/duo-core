package usecase

import (
	"github.com/lfcamarati/duo-core/domain/client/entity"
)

func NewDeleteClientUseCase(repository entity.ClientRepository) *DeleteClientUseCase {
	return &DeleteClientUseCase{repository}
}

type DeleteClientInput struct {
	ID int64
}

type DeleteClientOutput struct{}

type DeleteClientUseCase struct {
	Repository entity.ClientRepository
}

func (uc *DeleteClientUseCase) Execute(input DeleteClientInput) (*DeleteClientOutput, error) {
	err := uc.Repository.Delete(input.ID)

	if err != nil {
		return nil, err
	}

	return &DeleteClientOutput{}, nil
}
