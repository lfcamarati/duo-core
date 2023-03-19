package service

import (
	serviceInfra "github.com/lfcamarati/duo-core/infra/domain/service"
)

func NewDeleteServiceUseCase(factory serviceInfra.ServiceRepositoryFactory) DeleteServiceUseCase {
	return DeleteServiceUseCase{factory}
}

type DeleteServiceInput struct {
	ID int64
}

type DeleteServiceOutput struct{}

type DeleteServiceUseCase struct {
	NewRepository serviceInfra.ServiceRepositoryFactory
}

func (uc *DeleteServiceUseCase) Execute(input DeleteServiceInput) (*DeleteServiceOutput, error) {
	repository := uc.NewRepository()

	if err := repository.Begin(); err != nil {
		return nil, err
	}
	defer repository.Rollback()

	err := repository.Delete(input.ID)

	if err != nil {
		return nil, err
	}

	if err := repository.Commit(); err != nil {
		return nil, err
	}

	return &DeleteServiceOutput{}, nil
}
