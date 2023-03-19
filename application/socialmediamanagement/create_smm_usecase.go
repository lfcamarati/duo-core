package socialmediamanagement

import (
	ssmInfra "github.com/lfcamarati/duo-core/infra/domain/socialmediamanagement"
)

func NewCreateServiceUsecase(factory ssmInfra.SocialMediaManagementRepositoryFactory) CreateSocialMediaManagementService {
	return CreateSocialMediaManagementService{factory}
}

type CreateServiceUsecaseInput struct {
	Name        string
	Type        string
	Description string
	Price       float64
}

type CreateServiceUsecaseOutput struct {
	ID *int64
}

type CreateSocialMediaManagementService struct {
	NewRepository ssmInfra.SocialMediaManagementRepositoryFactory
}

func (s *CreateSocialMediaManagementService) Execute(input *CreateServiceUsecaseInput) (*CreateServiceUsecaseOutput, error) {
	return nil, nil
}
