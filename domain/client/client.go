package client

import (
	"errors"
)

var (
	ErrInvalidName       = errors.New("nome deve ser informado")
	ErrInvalidType       = errors.New("tipo do cliente deve ser informado")
	ErrInvalidCpf        = errors.New("cpf deve ser informado")
	ErrInvalidCnpj       = errors.New("cnpj deve ser informado")
	ErrInvalidAddress    = errors.New("endereco deve ser informado")
	ErrInvalidEmail      = errors.New("email deve ser informado")
	ErrInvalidPhone      = errors.New("telefone deve ser informado")
	ErrInvalidClientType = errors.New("tipo de cliente inválido")
)

func NewClient(name *string, cpfCnpj *string, address *string, email *string, phone *string, clientType *string) (Client, error) {
	err := validate(name, cpfCnpj, address, email, phone, clientType)

	if err != nil {
		return Client{}, err
	}

	return Client{
		Id:      nil,
		Name:    *name,
		CpfCnpj: *cpfCnpj,
		Address: *address,
		Email:   *email,
		Phone:   *phone,
		Type:    ClientType(*clientType),
	}, nil
}

func NewClientType(typeName string) (*ClientType, error) {
	clientType := ClientType(typeName)

	if clientType != PF && clientType != PJ {
		return nil, ErrInvalidClientType
	}

	return &clientType, nil
}

func validate(name *string, cpfCnpj *string, address *string, email *string, phone *string, clientType *string) error {
	if name == nil && *name == "" {
		return ErrInvalidName
	}

	if clientType == nil || *clientType == "" {
		return ErrInvalidType
	}

	if cpfCnpj == nil || *cpfCnpj == "" {
		if ClientType(*clientType) == PF {
			return ErrInvalidCpf
		} else {
			return ErrInvalidCnpj
		}
	}

	if address == nil || *address == "" {
		return ErrInvalidAddress
	}

	if email == nil || *email == "" {
		return ErrInvalidEmail
	}

	if phone == nil || *phone == "" {
		return ErrInvalidPhone
	}

	return nil
}

type ClientId int64
type ClientType string

const (
	PF ClientType = "PF"
	PJ ClientType = "PJ"
)

type Client struct {
	Id      *ClientId
	Name    string
	CpfCnpj string
	Address string
	Email   string
	Phone   string
	Type    ClientType
}
