package usecase

import (
	"github.com/lfcamarati/duo-core/domain/service/infra/repository"
)

func NewGetAllServicesUseCase(factory repository.ServiceRepositoryFactory) GetAllServicesUseCase {
	return GetAllServicesUseCase{factory}
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
	NewRepository repository.ServiceRepositoryFactory
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
			Title:       service.Title,
			Description: service.Description,
			Price:       service.Price,
		}

		servicesOutput = append(servicesOutput, serviceOutput)
	}

	return &GetAllServicesUseCaseOutput{Data: servicesOutput}, nil
}
