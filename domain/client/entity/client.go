package entity

func NewClientPf(name string, cpf string, address string, email string, phone string, clientType string) *ClientPf {
	return &ClientPf{
		Client{
			nil,
			address,
			email,
			phone,
			clientType,
		},
		name,
		cpf,
	}
}

func NewClientPj(corporateName string, cnpj string, address string, email string, phone string, clientType string) *ClientPj {
	return &ClientPj{
		Client{
			nil,
			address,
			email,
			phone,
			clientType,
		},
		corporateName,
		cnpj,
	}
}

type Client struct {
	ID      *int32 `json:"id"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Type    string `json:"type"`
}

func (c *Client) IsPf() bool {
	return c.Type == "PF"
}

type ClientPf struct {
	Client
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}

type ClientPj struct {
	Client
	CorporateName string `json:"corporateName"`
	Cnpj          string `json:"cnpj"`
}

type ClientSearch struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ClientRepository interface {
	SavePf(client *ClientPf) (*int64, error)
	SavePj(client *ClientPj) (*int64, error)
	GetAll() ([]ClientSearch, error)
	GetById(id int64) (*Client, error)
	Delete(id int64) error
}
