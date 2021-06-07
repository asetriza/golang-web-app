package service

import (
	"context"
	"errors"

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

var AccessDenied = errors.New("access denied")

func (ts *TodoService) Update(ctx context.Context, todo common.Todo) error {
	userID, err := ts.getUserIDFromTodo(ctx, todo.ID)
	if err != nil {
		return err
	}

	if !todo.IsOwner(userID) {
		return AccessDenied
	}

	return ts.Repository.Update(ctx, todo)
}

func (ts *TodoService) getUserIDFromTodo(ctx context.Context, todoID int) (int, error) {
	userID, err := ts.Repository.GetUserIDFromTodo(ctx, todoID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (ts *TodoService) Delete(ctx context.Context, userID, todoID int) error {
	userIDFromTodo, err := ts.getUserIDFromTodo(ctx, todoID)
	if err != nil {
		return err
	}

	if userIDFromTodo != userID {
		return AccessDenied
	}

	return ts.Repository.Delete(ctx, userID, todoID)
}
