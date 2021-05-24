package service

import (
	"context"
	"golang-web-app/internal/common"
	"golang-web-app/internal/repository"
	"golang-web-app/pkg/auth"
)

type Authorization interface {
	CreateUser(ctx context.Context, user common.User) (int, error)
	GenerateCredentials(ctx context.Context, username, password string) (Credentials, error)
	RefreshCredentials(ctx context.Context, token, refreshToken string) (Credentials, error)
}

type Todo interface {
	Create(ctx context.Context, todo common.Todo) (int, error)
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
