package usecase

import (
	"github.com/lfcamarati/duo-core/domain/service/entity"
	"github.com/lfcamarati/duo-core/domain/service/infra/repository"
)

func NewCreateServiceUsecase(factory repository.ServiceRepositoryFactory) CreateServiceUsecase {
	return CreateServiceUsecase{factory}
}

type CreateServiceUsecaseInput struct {
	Title       string
	Description string
	Price       float64
}

type CreateServiceUsecaseOutput struct {
	ID *int64
}

type CreateServiceUsecase struct {
	NewRepository repository.ServiceRepositoryFactory
}

func (uc *CreateServiceUsecase) Execute(input *CreateServiceUsecaseInput) (*CreateServiceUsecaseOutput, error) {
	repository := uc.NewRepository()

	if err := repository.Begin(); err != nil {
		return nil, err
	}
	defer repository.Rollback()

	service := entity.NewService(input.Title, &input.Description, input.Price)
	ID, err := repository.Save(service)

	if err != nil {
		return nil, err
	}

	if err := repository.Commit(); err != nil {
		return nil, err
	}

	return &CreateServiceUsecaseOutput{ID}, nil
}
