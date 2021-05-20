package service

import (
	"golang-web-app/internal/common"
	"golang-web-app/internal/repository"
)

type User interface {
	Create(user common.User) (common.User, error)
}

type Authorization interface {
	CreateUser(user common.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Dependencies struct {
	Repository *repository.Repository
}

type Service struct {
	Authorization Authorization
	User          User
}

func NewService(deps Dependencies) *Service {
	return &Service{
		Authorization: NewAuthorizationService(deps.Repository.Authorization),
		User:          NewUserService(deps.Repository.User),
	}
}
