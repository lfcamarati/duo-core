package service

import (
	serviceInfra "github.com/lfcamarati/duo-core/infra/domain/service"
)

func NewGetAllServicesUseCase(factory serviceInfra.ServiceRepositoryFactory) GetAllServicesUseCase {
	return GetAllServicesUseCase{factory}
}

type GetAllServicesUseCaseInput struct{}

type serviceOutput struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type GetAllServicesUseCaseOutput struct {
	Data []serviceOutput `json:"data"`
}

type GetAllServicesUseCase struct {
	NewRepository serviceInfra.ServiceRepositoryFactory
}

func (uc *GetAllServicesUseCase) Execute(input GetAllServicesUseCaseInput) (*GetAllServicesUseCaseOutput, error) {
	repository := uc.NewRepository()
	services, err := repository.GetAll()

	if err != nil {
		return nil, err
	}

	servicesOutput := make([]serviceOutput, 0)

	for _, service := range services {
		serviceOutput := serviceOutput{
			ID:          *service.ID,
			Name:        service.Name,
			Type:        string(service.Type),
			Description: service.Description,
			Price:       service.Price,
		}

		servicesOutput = append(servicesOutput, serviceOutput)
	}

	return &GetAllServicesUseCaseOutput{Data: servicesOutput}, nil
}
