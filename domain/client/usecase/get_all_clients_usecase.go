package usecase

import (
	"github.com/lfcamarati/duo-core/domain/client/entity"
)

func NewGetAllClientsUseCase(repository entity.ClientRepository) *GetAllClientsUseCase {
	return &GetAllClientsUseCase{repository}
}

type GetAllClientsInput struct{}

type GetAllClientsOutput struct {
	Data *[]entity.ClientSearch `json:"data"`
}

type GetAllClientsUseCase struct {
	Repository entity.ClientRepository
}

func (uc *GetAllClientsUseCase) Execute(input GetAllClientsInput) (*GetAllClientsOutput, error) {
	clients, err := uc.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return &GetAllClientsOutput{&clients}, nil
}
