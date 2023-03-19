package socialmediamanagement

import (
	ssmInfra "github.com/lfcamarati/duo-core/infra/domain/socialmediamanagement"
)

func NewGetServiceByIdUseCase(factory ssmInfra.SocialMediaManagementRepositoryFactory) GetServiceByIdUseCase {
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
	NewRepository ssmInfra.SocialMediaManagementRepositoryFactory
}

func (uc *GetServiceByIdUseCase) Execute(input GetServiceByIdUseCaseInput) (*GetServiceByIdUseCaseOutput, error) {
	return nil, nil
}
