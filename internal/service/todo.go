package service

import (
	"context"

	"github.com/asetriza/golang-web-app/internal/common"
	"github.com/asetriza/golang-web-app/internal/repository"
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

func (ts *TodoService) Get(ctx context.Context, userID, todoID int) (common.Todo, error) {
	return ts.Repository.Get(ctx, userID, todoID)
}

func (ts *TodoService) GetAll(ctx context.Context, userID int, pagination common.Pagination) ([]common.Todo, error) {
	return ts.Repository.GetAll(ctx, userID, pagination.CalculateOffset())
}

func (ts *TodoService) Update(ctx context.Context, todo common.Todo) error {
	return ts.Repository.Update(ctx, todo)
}

func (ts *TodoService) Delete(ctx context.Context, todoID int) error {
	return ts.Repository.Delete(ctx, todoID)
}
