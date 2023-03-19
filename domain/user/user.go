package user

func NewUser(name string, username string, password string) User {
	return User{
		Name:     name,
		Username: username,
		Password: password,
	}
}

type User struct {
	ID       *int64
	Name     string
	Username string
	Password string
}
