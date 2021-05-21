package service

import (
	"golang-web-app/internal/common"
	"golang-web-app/internal/repository"
	"golang-web-app/pkg/auth"
)

type User interface {
	Create(user common.User) (common.User, error)
}

type Authorization interface {
	CreateUser(user common.User) (int, error)
	GenerateCredentials(username, password string) (Credentials, error)
	RefreshCredentials(token, refreshToken string) (Credentials, error)
}

type Dependencies struct {
	Repository   *repository.Repository
	TokenManager auth.TokenManager
	PasswordSalt string
}

type Service struct {
	Authorization Authorization
	User          User
}

func NewService(deps Dependencies) *Service {
	return &Service{
		Authorization: NewAuthorizationService(deps.Repository.Authorization, deps.TokenManager, deps.PasswordSalt),
		User:          NewUserService(deps.Repository.User),
	}
}
