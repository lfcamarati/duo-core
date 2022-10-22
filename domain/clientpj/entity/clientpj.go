package entity

func NewClientPj(corporateName string, cnpj string, address string, email string, phone string, clientType string) *ClientPj {
	return &ClientPj{
		nil,
		address,
		email,
		phone,
		clientType,
		corporateName,
		cnpj,
	}
}

type ClientPj struct {
	ID            *int32 `json:"id"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Type          string `json:"type"`
	CorporateName string `json:"corporateName"`
	Cnpj          string `json:"cnpj"`
}

type ClientPjRepository interface {
	Save(client *ClientPj) (*int64, error)
	GetAll() ([]ClientPj, error)
	GetById(id int64) (*ClientPj, error)
	Delete(id int64) error
}
