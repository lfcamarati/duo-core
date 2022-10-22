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
	ID      *int64
	Type    string
	Name    string
	Cpf     string
	Address string
	Email   string
	Phone   string
}

type ClientPfRepository interface {
	Save(client ClientPf) (*int64, error)
	GetAll() ([]ClientPf, error)
	GetById(id int64) (*ClientPf, error)
	Delete(id int64) error
}
