package client

import (
	clientDomain "github.com/lfcamarati/duo-core/domain/client"
)

func newClientResource(client clientDomain.Client) ClientResource {
	return ClientResource{
		Id:      int64(*client.Id),
		Name:    client.Name,
		CpfCnpj: client.CpfCnpj,
		Address: client.Address,
		Email:   client.Email,
		Phone:   client.Phone,
		Type:    string(client.Type),
	}
}

func newClientListResource(clients []clientDomain.Client) ClientListResource {
	clientListResource := make([]ClientResource, 0)

	for _, client := range clients {
		clientListResource = append(clientListResource, newClientResource(client))
	}

	return ClientListResource{clientListResource}
}

type ClientResource struct {
	Id      int64  `json:"id"`
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
