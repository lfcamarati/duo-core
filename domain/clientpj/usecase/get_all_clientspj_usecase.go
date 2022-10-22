package usecase

import (
	clientPj "github.com/lfcamarati/duo-core/domain/clientpj/entity"
)

func NewGetAllClientsPjUseCase(clientPjRepository clientPj.ClientPjRepository) GetAllClientsPjUseCase {
	return GetAllClientsPjUseCase{clientPjRepository}
}

type GetAllClientsPjUseCaseInput struct{}

type GetAllClientsPjUseCaseOutput struct {
	Data []clientPj.ClientPj `json:"data"`
}

type GetAllClientsPjUseCase struct {
	Repository clientPj.ClientPjRepository
}

func (uc *GetAllClientsPjUseCase) Execute(input GetAllClientsPjUseCaseInput) (*GetAllClientsPjUseCaseOutput, error) {
	clientsPj, err := uc.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return &GetAllClientsPjUseCaseOutput{Data: clientsPj}, nil
}
