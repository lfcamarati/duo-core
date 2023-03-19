package socialmediamanagement

import (
	ssmInfra "github.com/lfcamarati/duo-core/infra/domain/socialmediamanagement"
)

func NewUpdateServiceUsecase(factory ssmInfra.SocialMediaManagementRepositoryFactory) *UpdateServiceUsecase {
	return &UpdateServiceUsecase{factory}
}

type UpdateServiceUsecaseInput struct {
	ID          int64
	Name        string
	Description string
	Price       float64
}

type UpdateServiceUsecaseOutput struct{}

type UpdateServiceUsecase struct {
	NewRepository ssmInfra.SocialMediaManagementRepositoryFactory
}

func (s *UpdateServiceUsecase) Execute(input UpdateServiceUsecaseInput) (*UpdateServiceUsecaseOutput, error) {
	return nil, nil
}
