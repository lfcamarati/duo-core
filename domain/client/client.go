package client

type ClientType string

func NewClient(name string, cpfCnpj string, address string, email string, phone string, clientType ClientType) Client {
	return Client{
		ID:      nil,
		Name:    name,
		CpfCnpj: cpfCnpj,
		Address: address,
		Email:   email,
		Phone:   phone,
		Type:    clientType,
	}
}

const (
	PF ClientType = "PF"
	PJ ClientType = "PJ"
)

type Client struct {
	ID      *int64
	Name    string
	CpfCnpj string
	Address string
	Email   string
	Phone   string
	Type    ClientType
}
