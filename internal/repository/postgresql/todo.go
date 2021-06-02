package postgesql

import (
	"context"

	"github.com/asetriza/golang-web-app/internal/common"

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
	rows, err := tr.db.NamedQueryContext(ctx,
		`insert into todos (
			user_id,
			name,
			description,
			notify_date,
			done)
		values (
			:user_id,
			:name,
			:description,
			:notify_date,
			:done)
		returning
			id;`, todo)
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
	var todo common.Todo
	err := tr.db.GetContext(ctx, &todo,
		`select
			id,
			user_id,
			name,
			description,
			notify_date,
			done
		from
			todos
		where
			id = $1;`, todoID)
	if err != nil {
		return common.Todo{}, err
	}

	return todo, nil
}

func (tr *TodoRepository) GetAll(ctx context.Context, userID int) ([]common.Todo, error) {
	var todos []common.Todo
	err := tr.db.SelectContext(ctx, &todos,
		`select
			id,
			user_id,
			name,
			description,
			notify_date,
			done
		from
			todos
		where
			user_id = $1;`, userID)
	if err != nil {
		return []common.Todo{}, err
	}

	return todos, nil
}

func (tr *TodoRepository) Update(ctx context.Context, todo common.Todo) error {
	return nil
}

func (tr *TodoRepository) Delete(ctx context.Context, todoID int) error {
	rows, err := tr.db.QueryContext(ctx,
		`delete from todos where id = $1;`, todoID)
	if err != nil {
		return err
	}

	rows.Close()

	return nil
}
