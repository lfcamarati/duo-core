package usecase

import (
	"github.com/lfcamarati/duo-core/domain/clientpf/infra/repository"
)

func NewUpdateClientPfUseCase(repository repository.ClientPfRepositoryFactory) *UpdateClientPfUseCase {
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
	NewRepository repository.ClientPfRepositoryFactory
}

func (uc *UpdateClientPfUseCase) Execute(input *UpdateClientPfUsecaseInput) (*UpdateClientPfUsecaseOutput, error) {
	repository := uc.NewRepository()
	clientPf, err := repository.GetById(input.ID)

	if err != nil {
		return nil, err
	}

	clientPf.Name = input.Name
	clientPf.Cpf = input.Cpf
	clientPf.Address = input.Address
	clientPf.Email = input.Email
	clientPf.Phone = input.Phone

	repository.Begin()
	err = repository.Update(*clientPf)

	if err != nil {
		repository.Rollback()
		return nil, err
	}

	repository.Commit()
	return &UpdateClientPfUsecaseOutput{}, nil
}
