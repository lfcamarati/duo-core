package entity

func NewUser(nome string, username string, password string) User {
	return User{
		Nome:     nome,
		Username: username,
		Password: password,
	}
}

type User struct {
	ID       *int64
	Nome     string
	Username string
	Password string
}

type UserRepository interface {
}
