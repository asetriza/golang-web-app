package service

import (
	"golang-web-app/internal/common"
	"golang-web-app/internal/repository"
)

type UserService struct {
	Repository repository.User
}

func NewUserService(ru repository.User) *UserService {
	return &UserService{
		Repository: ru,
	}
}

func (us *UserService) Create(common.User) (common.User, error) {
	return common.User{}, nil
}
