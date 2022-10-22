package usecase

import (
	"github.com/lfcamarati/duo-core/domain/clientpj/entity"
)

func NewUpdateClientPjUsecase(repository entity.ClientPjRepository) *UpdateClientPjUsecase {
	return &UpdateClientPjUsecase{repository}
}

type UpdateClientPjUsecaseInput struct {
	ID            int64
	CorporateName string
	Cnpj          string
	Address       string
	Email         string
	Phone         string
}

type UpdateClientPjUsecaseOutput struct{}

type UpdateClientPjUsecase struct {
	Repository entity.ClientPjRepository
}

func (uc *UpdateClientPjUsecase) Execute(input *UpdateClientPjUsecaseInput) (*UpdateClientPjUsecaseOutput, error) {
	clientPj, err := uc.Repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	clientPj.CorporateName = input.CorporateName
	clientPj.Cnpj = input.Cnpj
	clientPj.Address = input.Address
	clientPj.Email = input.Email
	clientPj.Phone = input.Phone

	err = uc.Repository.Update(*clientPj)

	if err != nil {
		return nil, err
	}

	return &UpdateClientPjUsecaseOutput{}, nil
}
