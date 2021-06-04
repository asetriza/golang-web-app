package service

import (
	"context"

	"github.com/asetriza/golang-web-app/internal/common"
	"github.com/asetriza/golang-web-app/internal/repository"
	"github.com/asetriza/golang-web-app/pkg/auth"
)

type Authorization interface {
	CreateUser(ctx context.Context, user common.User, userIP string) (Credentials, error)
	CreateCredentials(ctx context.Context, username, password, userIP string) (Credentials, error)
	RefreshCredentials(ctx context.Context, token, refreshToken, userIP string) (Credentials, error)
}

type Todo interface {
	Create(ctx context.Context, todo common.Todo) (int, error)
	Get(ctx context.Context, userID, todoID int) (common.Todo, error)
	GetAll(ctx context.Context, userID int, pagination common.Pagination) ([]common.Todo, error)
	Update(ctx context.Context, todo common.Todo) error
	Delete(ctx context.Context, todoID int) error
}

type Dependencies struct {
	Repository   *repository.Repository
	TokenManager auth.TokenManager
	PasswordSalt string
}

type Service struct {
	Authorization Authorization
	Todo          Todo
}

func NewService(deps Dependencies) *Service {
	return &Service{
		Authorization: NewAuthorizationService(deps.Repository.Authorization, deps.TokenManager, deps.PasswordSalt),
		Todo:          NewTodoService(deps.Repository.Todo),
	}
}
