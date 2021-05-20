package postgesql

import (
	"golang-web-app/internal/common"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(common.User) (common.User, error) {
	return common.User{}, nil
}
