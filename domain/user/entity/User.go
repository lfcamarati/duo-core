package entity

import (
	"context"

	"github.com/lfcamarati/duo-core/infra/database"
)

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

type UserRepository interface {
	database.Transactor

	Save(ctx context.Context, user User) (*int64, error)
	FindByUsername(username string) (*User, error)
}
