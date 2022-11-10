package usecase

import (
	"github.com/lfcamarati/duo-core/domain/clientpj/infra/repository"
)

func NewUpdateClientPjUsecase(factory repository.ClientPjRepositoryFactory) *UpdateClientPjUsecase {
	return &UpdateClientPjUsecase{factory}
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
	NewRepository repository.ClientPjRepositoryFactory
}

func (uc *UpdateClientPjUsecase) Execute(input *UpdateClientPjUsecaseInput) (*UpdateClientPjUsecaseOutput, error) {
	repository := uc.NewRepository()
	clientPj, err := repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	clientPj.CorporateName = input.CorporateName
	clientPj.Cnpj = input.Cnpj
	clientPj.Address = input.Address
	clientPj.Email = input.Email
	clientPj.Phone = input.Phone

	repository.Begin()
	err = repository.Update(*clientPj)

	if err != nil {
		repository.Rollback()
		return nil, err
	}

	repository.Commit()
	return &UpdateClientPjUsecaseOutput{}, nil
}
