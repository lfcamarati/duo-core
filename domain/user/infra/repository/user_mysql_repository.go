package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lfcamarati/duo-core/domain/user/entity"
)

func NewUserRepository(db *sql.DB) entity.UserRepository {
	return UserMysqlRepository{db}
}

type UserMysqlRepository struct {
	db *sql.DB
}
