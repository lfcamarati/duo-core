package entity

func NewClientPf(name string, cpf string, address string, email string, phone string) ClientPf {
	return ClientPf{
		ID:      nil,
		Type:    "PF",
		Name:    name,
		Cpf:     cpf,
		Address: address,
		Email:   email,
		Phone:   phone,
	}
}

type ClientPf struct {
	ID      *int64 `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Cpf     string `json:"cpf"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

type ClientPfRepository interface {
	Save(client ClientPf) (*int64, error)
	GetAll() ([]ClientPf, error)
	GetById(id int64) (*ClientPf, error)
	Delete(id int64) error
}
