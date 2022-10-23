package usecase

import (
	"github.com/lfcamarati/duo-core/domain/clientpf/entity"
)

func NewCreateClientPfUseCase(repository entity.ClientPfRepository) *CreateClientPfUseCase {
	return &CreateClientPfUseCase{repository}
}

type CreateClientPfUsecaseInput struct {
	Name    string
	Cpf     string
	Address string
	Email   string
	Phone   string
}

type CreateClientPfUsecaseOutput struct {
	ID *int64
}

type CreateClientPfUseCase struct {
	Repository entity.ClientPfRepository
}

func (uc *CreateClientPfUseCase) Execute(input *CreateClientPfUsecaseInput) (*CreateClientPfUsecaseOutput, error) {
	clientPf := entity.NewClientPf(input.Name, input.Cpf, input.Address, input.Email, input.Phone)
	ID, err := uc.Repository.Save(clientPf)

	if err != nil {
		return nil, err
	}

	return &CreateClientPfUsecaseOutput{ID}, nil
}
