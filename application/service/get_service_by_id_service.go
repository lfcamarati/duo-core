package service

import (
	serviceInfra "github.com/lfcamarati/duo-core/infra/domain/service"
)

func NewGetServiceByIdUseCase(factory serviceInfra.ServiceRepositoryFactory) GetServiceByIdUseCase {
	return GetServiceByIdUseCase{factory}
}

type GetServiceByIdUseCaseInput struct {
	ID int64
}

type GetServiceByIdUseCaseOutput struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type GetServiceByIdUseCase struct {
	NewRepository serviceInfra.ServiceRepositoryFactory
}

func (uc *GetServiceByIdUseCase) Execute(input GetServiceByIdUseCaseInput) (*GetServiceByIdUseCaseOutput, error) {
	repository := uc.NewRepository()
	service, err := repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	return &GetServiceByIdUseCaseOutput{
		ID:          int64(*service.ID),
		Name:        service.Name,
		Type:        string(service.Type),
		Description: service.Description,
		Price:       service.Price,
	}, nil
}
