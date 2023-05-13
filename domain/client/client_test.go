package client

import (
	"testing"
)

func TestCreateValidClient(t *testing.T) {
	name := "client"
	cpfCnpj := "12345678900"
	address := "Address Abc"
	email := "email@email.com"
	phone := "11222223333"
	clientType := "PF"

	_, err := NewClient(&name, &cpfCnpj, &address, &email, &phone, &clientType)

	if err != nil {
		t.Errorf("Error ocurred = %s", err)
	}
}
