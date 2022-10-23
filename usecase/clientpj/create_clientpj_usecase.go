package usecase

import (
	"github.com/lfcamarati/duo-core/domain/clientpj/entity"
)

func NewCreateClientPjUsecase(repository entity.ClientPjRepository) *CreateClientPjUseCase {
	return &CreateClientPjUseCase{repository}
}

type CreateClientPjInput struct {
	Type          string
	CorporateName string
	Cnpj          string
	Address       string
	Email         string
	Phone         string
}

type CreateClientPjOutput struct {
	ID *int64
}

type CreateClientPjUseCase struct {
	Repository entity.ClientPjRepository
}

func (uc *CreateClientPjUseCase) Execute(input *CreateClientPjInput) (*CreateClientPjOutput, error) {
	clientPj := entity.NewClientPj(input.CorporateName, input.Cnpj, input.Address, input.Email, input.Phone)
	ID, err := uc.Repository.Save(clientPj)

	if err != nil {
		return nil, err
	}

	return &CreateClientPjOutput{ID}, nil
}
