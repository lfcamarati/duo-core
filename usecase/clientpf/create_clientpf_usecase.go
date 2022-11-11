package usecase

import (
	"errors"

	"github.com/lfcamarati/duo-core/domain/clientpf/entity"
	"github.com/lfcamarati/duo-core/domain/clientpf/infra/repository"
)

func NewCreateClientPfUseCase(factory repository.ClientPfRepositoryFactory) *CreateClientPfUseCase {
	return &CreateClientPfUseCase{factory}
}

type CreateClientPfUsecaseInput struct {
	Name    *string
	Cpf     *string
	Address *string
	Email   *string
	Phone   *string
}

func (i *CreateClientPfUsecaseInput) Validate() error {
	if i.Name == nil || *i.Name == "" {
		return errors.New("nome deve ser informado")
	}

	if i.Cpf == nil || *i.Cpf == "" {
		return errors.New("cpf deve ser informado")
	}

	if i.Address == nil || *i.Address == "" {
		return errors.New("endereço deve ser informado")
	}

	if i.Email == nil || *i.Email == "" {
		return errors.New("email deve ser informado")
	}

	if i.Phone == nil || *i.Phone == "" {
		return errors.New("telefone deve ser informado")
	}

	return nil
}

type CreateClientPfUsecaseOutput struct {
	ID *int64 `json:"id"`
}

type CreateClientPfUseCase struct {
	NewRepository repository.ClientPfRepositoryFactory
}

func (uc *CreateClientPfUseCase) Execute(input *CreateClientPfUsecaseInput) (*CreateClientPfUsecaseOutput, error) {
	err := input.Validate()

	if err != nil {
		return nil, err
	}

	repository := uc.NewRepository()

	if err := repository.Begin(); err != nil {
		return nil, err
	}
	defer repository.Rollback()

	clientPf := entity.NewClientPf(*input.Name, *input.Cpf, *input.Address, *input.Email, *input.Phone)
	ID, err := repository.Save(clientPf)

	if err != nil {
		return nil, err
	}

	if err := repository.Commit(); err != nil {
		return nil, err
	}

	return &CreateClientPfUsecaseOutput{ID}, nil
}
