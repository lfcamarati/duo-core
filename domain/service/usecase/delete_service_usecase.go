package usecase

import (
	"github.com/lfcamarati/duo-core/domain/service/entity"
)

func NewDeleteServiceUseCase(repository entity.ServiceRepository) DeleteServiceUseCase {
	return DeleteServiceUseCase{repository}
}

type DeleteServiceInput struct {
	ID int64
}

type DeleteServiceOutput struct{}

type DeleteServiceUseCase struct {
	Repository entity.ServiceRepository
}

func (uc *DeleteServiceUseCase) Execute(input DeleteServiceInput) (*DeleteServiceOutput, error) {
	err := uc.Repository.Delete(input.ID)

	if err != nil {
		return nil, err
	}

	return &DeleteServiceOutput{}, nil
}
