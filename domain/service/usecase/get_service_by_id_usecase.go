package usecase

import (
	"github.com/lfcamarati/duo-core/domain/service/entity"
)

func NewGetServiceByIdUseCase(clientPfRepository entity.ServiceRepository) GetServiceByIdUseCase {
	return GetServiceByIdUseCase{clientPfRepository}
}

type GetServiceByIdUseCaseInput struct {
	ID int64
}

type GetServiceByIdUseCaseOutput struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Price       float64 `json:"price"`
}

type GetServiceByIdUseCase struct {
	Repository entity.ServiceRepository
}

func (uc *GetServiceByIdUseCase) Execute(input GetServiceByIdUseCaseInput) (*GetServiceByIdUseCaseOutput, error) {
	service, err := uc.Repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	return &GetServiceByIdUseCaseOutput{
		ID:          *service.ID,
		Title:       service.Title,
		Description: service.Description,
		Price:       service.Price,
	}, nil
}
