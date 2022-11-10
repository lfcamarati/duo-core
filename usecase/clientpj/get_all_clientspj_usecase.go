package usecase

import (
	clientPj "github.com/lfcamarati/duo-core/domain/clientpj/entity"
	"github.com/lfcamarati/duo-core/domain/clientpj/infra/repository"
)

func NewGetAllClientsPjUseCase(factory repository.ClientPjRepositoryFactory) GetAllClientsPjUseCase {
	return GetAllClientsPjUseCase{factory}
}

type GetAllClientsPjUseCaseInput struct{}

type GetAllClientsPjUseCaseOutput struct {
	Data []clientPj.ClientPj `json:"data"`
}

type GetAllClientsPjUseCase struct {
	NewRepository repository.ClientPjRepositoryFactory
}

func (uc *GetAllClientsPjUseCase) Execute(input GetAllClientsPjUseCaseInput) (*GetAllClientsPjUseCaseOutput, error) {
	repository := uc.NewRepository()
	clientsPj, err := repository.GetAll()

	if err != nil {
		return nil, err
	}

	return &GetAllClientsPjUseCaseOutput{Data: clientsPj}, nil
}
