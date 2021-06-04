package repository

import (
	"context"

	"github.com/asetriza/golang-web-app/internal/common"
	psqlrepo "github.com/asetriza/golang-web-app/internal/repository/postgresql"

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
	Get(ctx context.Context, todoID int) (common.Todo, error)
	GetAll(ctx context.Context, userID int, pagination common.Pagination) ([]common.Todo, error)
	Update(ctx context.Context, todo common.Todo) error
	Delete(ctx context.Context, todoID int) error
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
