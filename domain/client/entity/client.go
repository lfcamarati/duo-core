package entity

type Client struct {
	ID   *int64
	Name string
	Type string
}

type ClientRepository interface {
	GetAll() ([]Client, error)
	GetById(id int64) (*Client, error)
}
