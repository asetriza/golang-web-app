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
	rows, err := tr.db.NamedQueryContext(ctx, `
		insert into todos (
			user_id,
			name,
			description,
			notify_date,
			done
		)
		values (
			:user_id,
			:name,
			:description,
			:notify_date,
			:done
		)
		returning
			id; `, todo)
	if err != nil {
		return 0, err
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (tr *TodoRepository) Get(ctx context.Context, todoID int) (common.Todo, error) {
	return common.Todo{}, nil
}

func (tr *TodoRepository) GetAll(ctx context.Context) ([]common.Todo, error) {
	return []common.Todo{}, nil
}

func (tr *TodoRepository) Update(ctx context.Context, todo common.Todo) (int, error) {
	return 0, nil
}

func (tr *TodoRepository) Delete(ctx context.Context, todoID int) (int, error) {
	return 0, nil
}
