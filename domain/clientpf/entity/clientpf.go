package entity

func NewClientPf(name string, cpf string, address string, email string, phone string, clientType string) *ClientPf {
	return &ClientPf{
		nil,
		address,
		email,
		phone,
		clientType,
		name,
		cpf,
	}
}

type ClientPf struct {
	ID      *int32 `json:"id"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Cpf     string `json:"cpf"`
}

type ClientPfRepository interface {
	Save(client *ClientPf) (*int64, error)
	GetAll() ([]ClientPf, error)
	GetById(id int64) (*ClientPf, error)
	Delete(id int64) error
}
