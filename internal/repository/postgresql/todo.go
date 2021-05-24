package postgesql

import (
	"context"
	"golang-web-app/internal/common"

	"github.com/jmoiron/sqlx"
)

type TodoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (tr *TodoRepository) Create(ctx context.Context, todo common.Todo) (int, error) {
	return 0, nil
}
