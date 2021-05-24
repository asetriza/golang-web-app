package repository

import (
	"context"
	"golang-web-app/internal/common"
	psqlrepo "golang-web-app/internal/repository/postgresql"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(ctx context.Context, user common.User) (int, error)
	GetUser(ctx context.Context, username, password string) (common.User, error)
	GetUserSession(ctx context.Context, userID int, refreshToken string) (common.UserSession, error)
	CreateUserSession(ctx context.Context, userID int, userIP, refreshToken string, freshTokenTTL int64) (int, error)
	UpdateUserSession(ctx context.Context, userID int, refreshToken string, refreshTokenTTL int64) (int, error)
}

type Todo interface {
	Create(ctx context.Context, todo common.Todo) (int, error)
}

type Repository struct {
	Authorization Authorization
	Todo          Todo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: psqlrepo.NewAuthorizationRepository(db),
		Todo:          psqlrepo.NewTodoRepository(db),
	}
}
