package repository

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/domain/user/entity"
	"github.com/lfcamarati/duo-core/infra/database"
)

func NewUserRepositoryFactory(db *sql.DB) UserRepositoryFactory {
	return func() entity.UserRepository {
		return &UserMysqlRepository{
			&database.GenericTransactor{Db: db},
		}
	}
}

type UserRepositoryFactory func() entity.UserRepository

type UserMysqlRepository struct {
	*database.GenericTransactor
}

var (
	ErrTransactionInProgress = errors.New("transaction already in progress")
	ErrTransactionNotStarted = errors.New("transaction not started")
)

func (r UserMysqlRepository) Save(ctx context.Context, user entity.User) (*int64, error) {
	stmt, err := r.Tx.Prepare("INSERT INTO user (name, username, password) VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	rs, err := stmt.Exec(user.Name, user.Username, user.Password)

	if err != nil {
		return nil, err
	}

	id, err := rs.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}
