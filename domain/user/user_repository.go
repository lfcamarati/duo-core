package user

import (
	"context"

	"github.com/lfcamarati/duo-core/infra/database"
)

type UserRepository interface {
	database.Transactor

	Save(ctx context.Context, user User) (*int64, error)
	FindByUsername(username string) (*User, error)
}
