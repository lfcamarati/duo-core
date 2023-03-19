package socialmediamanagement

import (
	ssmInfra "github.com/lfcamarati/duo-core/infra/domain/socialmediamanagement"
)

func NewGetAllServicesUseCase(factory ssmInfra.SocialMediaManagementRepositoryFactory) GetAllServicesUseCase {
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
	NewRepository ssmInfra.SocialMediaManagementRepositoryFactory
}

func (s *GetAllServicesUseCase) Execute(input GetAllServicesUseCaseInput) (*GetAllServicesUseCaseOutput, error) {
	return nil, nil
}
