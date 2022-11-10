package usecase

import (
	"github.com/lfcamarati/duo-core/domain/service/infra/repository"
)

func NewGetServiceByIdUseCase(factory repository.ServiceRepositoryFactory) GetServiceByIdUseCase {
	return GetServiceByIdUseCase{factory}
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
	NewRepository repository.ServiceRepositoryFactory
}

func (uc *GetServiceByIdUseCase) Execute(input GetServiceByIdUseCaseInput) (*GetServiceByIdUseCaseOutput, error) {
	repository := uc.NewRepository()
	service, err := repository.GetById(input.ID)

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
