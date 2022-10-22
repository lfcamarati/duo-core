package usecase

import (
	"github.com/lfcamarati/duo-core/domain/clientpj/entity"
)

func NewDeleteClientPjUseCase(repository entity.ClientPjRepository) *DeleteClientPjUseCase {
	return &DeleteClientPjUseCase{repository}
}

type DeleteClientPjInput struct {
	ID int64
}

type DeleteClientPjOutput struct{}

type DeleteClientPjUseCase struct {
	Repository entity.ClientPjRepository
}

func (uc *DeleteClientPjUseCase) Execute(input DeleteClientPjInput) (*DeleteClientPjOutput, error) {
	err := uc.Repository.Delete(input.ID)

	if err != nil {
		return nil, err
	}

	return &DeleteClientPjOutput{}, nil
}
