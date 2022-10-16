package usecase

import (
	"errors"

	"github.com/lfcamarati/duo-core/internal/client/domain"
)

func NewCreateClientUseCase(repository domain.ClientRepository) *CreateClientUseCase {
	return &CreateClientUseCase{repository}
}

type CreateClientInput struct {
	Type          string
	Name          *string
	Cpf           *string
	CorporateName *string
	Cnpj          *string
	Address       string
	Email         string
	Phone         string
}

func (c *CreateClientInput) IsPf() bool {
	return c.Type == "PF"
}

func (c *CreateClientInput) IsInvalid() bool {
	if c.IsPf() {
		return c.Name == nil || c.Cpf == nil
	}

	return c.CorporateName == nil || c.Cnpj == nil
}

type CreateClientOutput struct {
	ID *int64
}

type CreateClientUseCase struct {
	Repository domain.ClientRepository
}

func (uc *CreateClientUseCase) Execute(input CreateClientInput) (*CreateClientOutput, error) {
	if input.IsInvalid() {
		return nil, errors.New("dados informados são inválidos")
	}

	if input.IsPf() {
		return uc.createClientPf(input)
	}

	return uc.createClientPj(input)
}

func (uc *CreateClientUseCase) createClientPf(input CreateClientInput) (*CreateClientOutput, error) {
	clientPf := domain.NewClientPf(*input.Name, *input.Cpf, input.Address, input.Email, input.Phone, input.Type)
	ID, err := uc.Repository.SavePf(clientPf)

	if err != nil {
		return nil, err
	}

	return &CreateClientOutput{ID}, nil
}

func (uc *CreateClientUseCase) createClientPj(input CreateClientInput) (*CreateClientOutput, error) {
	clientPj := domain.NewClientPj(*input.CorporateName, *input.Cnpj, input.Address, input.Email, input.Phone, input.Type)
	ID, err := uc.Repository.SavePj(clientPj)

	if err != nil {
		return nil, err
	}

	return &CreateClientOutput{ID}, nil
}
