package usecase

import (
	clientPf "github.com/lfcamarati/duo-core/domain/clientpf/entity"
)

func NewGetAllClientsPfUseCase(clientPfRepository clientPf.ClientPfRepository) GetAllClientsPfUseCase {
	return GetAllClientsPfUseCase{clientPfRepository}
}

type GetAllClientsPfUseCaseInput struct{}

type GetAllClientsPfUseCaseOutput struct {
	Data []clientPf.ClientPf `json:"data"`
}

type GetAllClientsPfUseCase struct {
	Repository clientPf.ClientPfRepository
}

func (uc *GetAllClientsPfUseCase) Execute(input GetAllClientsPfUseCaseInput) (*GetAllClientsPfUseCaseOutput, error) {
	clientsPf, err := uc.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return &GetAllClientsPfUseCaseOutput{Data: clientsPf}, nil
}
