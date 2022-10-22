package usecase

import (
	"github.com/lfcamarati/duo-core/domain/service/entity"
)

func NewGetAllServicesUseCase(clientPjRepository entity.ServiceRepository) GetAllServicesUseCase {
	return GetAllServicesUseCase{clientPjRepository}
}

type GetAllServicesUseCaseInput struct{}

type serviceOutput struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Price       float64 `json:"price"`
}

type GetAllServicesUseCaseOutput struct {
	Data []serviceOutput `json:"data"`
}

type GetAllServicesUseCase struct {
	Repository entity.ServiceRepository
}

func (uc *GetAllServicesUseCase) Execute(input GetAllServicesUseCaseInput) (*GetAllServicesUseCaseOutput, error) {
	services, err := uc.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	servicesOutput := make([]serviceOutput, 0)

	for _, service := range services {
		serviceOutput := serviceOutput{
			ID:          *service.ID,
			Title:       service.Title,
			Description: service.Description,
			Price:       service.Price,
		}

		servicesOutput = append(servicesOutput, serviceOutput)
	}

	return &GetAllServicesUseCaseOutput{Data: servicesOutput}, nil
}
