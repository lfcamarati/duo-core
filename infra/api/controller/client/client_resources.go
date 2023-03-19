package client

import (
	clientDomain "github.com/lfcamarati/duo-core/domain/client"
)

func NewClientResource(client clientDomain.Client) ClientResource {
	return ClientResource{
		ID:      *client.ID,
		Name:    client.Name,
		CpfCnpj: client.CpfCnpj,
		Address: client.Address,
		Email:   client.Email,
		Phone:   client.Phone,
		Type:    string(client.Type),
	}
}

func NewClientListResource(clients []clientDomain.Client) ClientListResource {
	clientsResource := make([]ClientResource, 0)

	for _, c := range clients {
		clientsResource = append(clientsResource, NewClientResource(c))
	}

	return ClientListResource{clientsResource}
}

type ClientResource struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	CpfCnpj string `json:"cpfCnpj"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Type    string `json:"type"`
}

type ClientListResource struct {
	Data []ClientResource `json:"data"`
}
