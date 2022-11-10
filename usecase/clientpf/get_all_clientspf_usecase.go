package usecase

import (
	clientPf "github.com/lfcamarati/duo-core/domain/clientpf/entity"
	"github.com/lfcamarati/duo-core/domain/clientpf/infra/repository"
)

func NewGetAllClientsPfUseCase(factory repository.ClientPfRepositoryFactory) GetAllClientsPfUseCase {
	return GetAllClientsPfUseCase{factory}
}

type GetAllClientsPfUseCaseInput struct{}

type GetAllClientsPfUseCaseOutput struct {
	Data []clientPf.ClientPf `json:"data"`
}

type GetAllClientsPfUseCase struct {
	NewRepository repository.ClientPfRepositoryFactory
}

func (uc *GetAllClientsPfUseCase) Execute(input GetAllClientsPfUseCaseInput) (*GetAllClientsPfUseCaseOutput, error) {
	repository := uc.NewRepository()
	clientsPf, err := repository.GetAll()

	if err != nil {
		return nil, err
	}

	return &GetAllClientsPfUseCaseOutput{Data: clientsPf}, nil
}
