package socialmediamanagement

import (
	ssmInfra "github.com/lfcamarati/duo-core/infra/domain/socialmediamanagement"
)

func NewDeleteServiceUseCase(factory ssmInfra.SocialMediaManagementRepositoryFactory) DeleteServiceUseCase {
	return DeleteServiceUseCase{factory}
}

type DeleteServiceInput struct {
	ID int64
}

type DeleteServiceOutput struct{}

type DeleteServiceUseCase struct {
	NewRepository ssmInfra.SocialMediaManagementRepositoryFactory
}

func (s *DeleteServiceUseCase) Execute(input DeleteServiceInput) (*DeleteServiceOutput, error) {
	return nil, nil
}
