package usecase

import (
	"github.com/lfcamarati/duo-core/domain/service/infra/repository"
)

func NewDeleteServiceUseCase(factory repository.ServiceRepositoryFactory) DeleteServiceUseCase {
	return DeleteServiceUseCase{factory}
}

type DeleteServiceInput struct {
	ID int64
}

type DeleteServiceOutput struct{}

type DeleteServiceUseCase struct {
	NewRepository repository.ServiceRepositoryFactory
}

func (uc *DeleteServiceUseCase) Execute(input DeleteServiceInput) (*DeleteServiceOutput, error) {
	repository := uc.NewRepository()
	repository.Begin()
	err := repository.Delete(input.ID)

	if err != nil {
		repository.Rollback()
		return nil, err
	}

	repository.Commit()
	return &DeleteServiceOutput{}, nil
}
