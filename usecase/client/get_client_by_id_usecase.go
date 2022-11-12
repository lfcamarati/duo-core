package usecase

import (
	"github.com/lfcamarati/duo-core/domain/client/infra/repository"
)

func NewGetClientByIdUseCase(factory repository.ClientRepositoryFactory) GetClientByIdUseCase {
	return GetClientByIdUseCase{factory}
}

type GetClientByIdUseCaseInput struct {
	ID int64
}

type GetClientByIdUseCaseOutput struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetClientByIdUseCase struct {
	NewRepository repository.ClientRepositoryFactory
}

func (uc *GetClientByIdUseCase) Execute(input GetClientByIdUseCaseInput) (*GetClientByIdUseCaseOutput, error) {
	repository := uc.NewRepository()
	clientPf, err := repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	return &GetClientByIdUseCaseOutput{
		ID:   *clientPf.ID,
		Name: clientPf.Name,
		Type: clientPf.Type,
	}, nil
}
