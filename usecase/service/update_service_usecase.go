package usecase

import (
	"github.com/lfcamarati/duo-core/domain/service/infra/repository"
)

func NewUpdateServiceUsecase(factory repository.ServiceRepositoryFactory) *UpdateServiceUsecase {
	return &UpdateServiceUsecase{factory}
}

type UpdateServiceUsecaseInput struct {
	ID          int64
	Title       string
	Description string
	Price       float64
}

type UpdateServiceUsecaseOutput struct{}

type UpdateServiceUsecase struct {
	NewRepository repository.ServiceRepositoryFactory
}

func (uc *UpdateServiceUsecase) Execute(input UpdateServiceUsecaseInput) (*UpdateServiceUsecaseOutput, error) {
	repository := uc.NewRepository()
	service, err := repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	service.Title = input.Title
	service.Description = &input.Description
	service.Price = input.Price

	repository.Begin()
	err = repository.Update(*service)

	if err != nil {
		repository.Rollback()
		return nil, err
	}

	repository.Commit()
	return &UpdateServiceUsecaseOutput{}, nil
}
