package usecase

import (
	"github.com/lfcamarati/duo-core/domain/clientpf/entity"
)

func NewUpdateClientPfUseCase(repository entity.ClientPfRepository) *UpdateClientPfUseCase {
	return &UpdateClientPfUseCase{repository}
}

type UpdateClientPfUsecaseInput struct {
	ID      int64
	Name    string
	Cpf     string
	Address string
	Email   string
	Phone   string
}

type UpdateClientPfUsecaseOutput struct{}

type UpdateClientPfUseCase struct {
	Repository entity.ClientPfRepository
}

func (uc *UpdateClientPfUseCase) Execute(input *UpdateClientPfUsecaseInput) (*UpdateClientPfUsecaseOutput, error) {
	clientPf, err := uc.Repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	clientPf.Name = input.Name
	clientPf.Cpf = input.Cpf
	clientPf.Address = input.Address
	clientPf.Email = input.Email
	clientPf.Phone = input.Phone

	err = uc.Repository.Update(*clientPf)

	if err != nil {
		return nil, err
	}

	return &UpdateClientPfUsecaseOutput{}, nil
}
