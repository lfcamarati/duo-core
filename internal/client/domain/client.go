package domain

func NewClientPf(name string, cpf string, address string, email string, phone string) *ClientPf {
	return &ClientPf{
		Client{
			nil,
			address,
			email,
			phone,
		},
		name,
		cpf,
	}
}

func NewClientPj(corporateName string, cnpj string, address string, email string, phone string) *ClientPj {
	return &ClientPj{
		Client{
			nil,
			address,
			email,
			phone,
		},
		corporateName,
		cnpj,
	}
}

type Client struct {
	ID      *int32
	Address string
	Email   string
	Phone   string
}

type ClientPf struct {
	Client
	Name string
	Cpf  string
}

type ClientPj struct {
	Client
	CorporateName string
	Cnpj          string
}

type ClientRepository interface {
	SavePf(client *ClientPf) (*int64, error)
	SavePj(client *ClientPj) (*int64, error)
}
