package entity

func NewClientPj(corporateName string, cnpj string, address string, email string, phone string) ClientPj {
	return ClientPj{
		ID:            nil,
		Type:          "PJ",
		CorporateName: corporateName,
		Cnpj:          cnpj,
		Address:       address,
		Email:         email,
		Phone:         phone,
	}
}

type ClientPj struct {
	ID            *int64 `json:"id"`
	Type          string `json:"type"`
	CorporateName string `json:"corporateName"`
	Cnpj          string `json:"cnpj"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
}

type ClientPjRepository interface {
	Save(client ClientPj) (*int64, error)
	GetAll() ([]ClientPj, error)
	GetById(id int64) (*ClientPj, error)
	Delete(id int64) error
}
