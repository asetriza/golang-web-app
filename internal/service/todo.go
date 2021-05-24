package service

import (
	"context"
	"golang-web-app/internal/common"
	"golang-web-app/internal/repository"
)

type TodoService struct {
	Repository repository.Todo
}

func NewTodoService(ru repository.Todo) *TodoService {
	return &TodoService{
		Repository: ru,
	}
}

func (ts *TodoService) Create(ctx context.Context, todo common.Todo) (int, error) {
	return ts.Repository.Create(ctx, todo)
}
