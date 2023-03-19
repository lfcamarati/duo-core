package user

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/domain/user"
	"github.com/lfcamarati/duo-core/infra/database"
)

func NewUserRepositoryFactory(db *sql.DB) UserRepositoryFactory {
	return func() user.UserRepository {
		return &UserMysqlRepository{
			&database.GenericTransactor{Db: db},
		}
	}
}

type UserRepositoryFactory func() user.UserRepository

type UserMysqlRepository struct {
	*database.GenericTransactor
}

var (
	ErrTransactionInProgress = errors.New("transaction already in progress")
	ErrTransactionNotStarted = errors.New("transaction not started")
)

func (r UserMysqlRepository) Save(ctx context.Context, user user.User) (*int64, error) {
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

func (r UserMysqlRepository) FindByUsername(username string) (*user.User, error) {
	user := new(user.User)

	err := r.Db.QueryRow(`
		SELECT
			u.id,
			u.name,
			u.username,
			u.password
		FROM user u
		WHERE u.username = ?`, username).Scan(&user.ID, &user.Name, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}
