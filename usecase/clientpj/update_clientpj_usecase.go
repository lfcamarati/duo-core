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

	if err := repository.Begin(); err != nil {
		return nil, err
	}
	defer repository.Rollback()

	clientPj, err := repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	clientPj.CorporateName = input.CorporateName
	clientPj.Cnpj = input.Cnpj
	clientPj.Address = input.Address
	clientPj.Email = input.Email
	clientPj.Phone = input.Phone

	err = repository.Update(*clientPj)

	if err != nil {
		return nil, err
	}

	if err := repository.Commit(); err != nil {
		return nil, err
	}

	return &UpdateClientPjUsecaseOutput{}, nil
}
