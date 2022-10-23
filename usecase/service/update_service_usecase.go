package usecase

import (
	"github.com/lfcamarati/duo-core/domain/service/entity"
)

func NewUpdateServiceUsecase(repository entity.ServiceRepository) *UpdateServiceUsecase {
	return &UpdateServiceUsecase{repository}
}

type UpdateServiceUsecaseInput struct {
	ID          int64
	Title       string
	Description string
	Price       float64
}

type UpdateServiceUsecaseOutput struct{}

type UpdateServiceUsecase struct {
	Repository entity.ServiceRepository
}

func (uc *UpdateServiceUsecase) Execute(input UpdateServiceUsecaseInput) (*UpdateServiceUsecaseOutput, error) {
	service, err := uc.Repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	service.Title = input.Title
	service.Description = &input.Description
	service.Price = input.Price

	err = uc.Repository.Update(*service)

	if err != nil {
		return nil, err
	}

	return &UpdateServiceUsecaseOutput{}, nil
}
