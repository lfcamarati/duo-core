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
	service := entity.NewService(input.Title, &input.Description, input.Price)

	repository.Begin()
	ID, err := repository.Save(service)

	if err != nil {
		repository.Rollback()
		return nil, err
	}

	repository.Commit()
	return &CreateServiceUsecaseOutput{ID}, nil
}
