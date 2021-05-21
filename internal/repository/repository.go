package repository

import (
	"golang-web-app/internal/common"
	psqlrepo "golang-web-app/internal/repository/postgresql"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user common.User) (int, error)
	GetUser(username, password string) (common.User, error)
	GetUserSession(userID int, refreshToken string) (common.UserSession, error)
	CreateUserSession(userID int, refreshToken string, freshTokenTTL int64) (int, error)
	UpdateUserSession(userID int, refreshToken string, refreshTokenTTL int64) (int, error)
}

type User interface {
	Create(common.User) (common.User, error)
}

type Repository struct {
	Authorization Authorization
	User          User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: psqlrepo.NewAuthorizationRepository(db),
		User:          psqlrepo.NewUserRepository(db),
	}
}
