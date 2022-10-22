package usecase

import (
	"github.com/lfcamarati/duo-core/domain/service/entity"
)

func NewCreateServiceUsecase(repository entity.ServiceRepository) CreateServiceUsecase {
	return CreateServiceUsecase{repository}
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
	Repository entity.ServiceRepository
}

func (uc *CreateServiceUsecase) Execute(input *CreateServiceUsecaseInput) (*CreateServiceUsecaseOutput, error) {
	service := entity.NewService(input.Title, &input.Description, input.Price)
	ID, err := uc.Repository.Save(service)

	if err != nil {
		return nil, err
	}

	return &CreateServiceUsecaseOutput{ID}, nil
}
