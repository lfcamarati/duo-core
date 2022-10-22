package entity

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

type ClientRepository interface {
	Save(client Client) (*int64, error)
	GetById(id int64) (*Client, error)
	Delete(id int64) error
}
