package postgesql

import (
	"context"
	"errors"
	"fmt"

	"github.com/asetriza/golang-web-app/internal/common"
	"github.com/asetriza/golang-web-app/pkg/database/postgresql"

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

func (tr *TodoRepository) Get(ctx context.Context, userID, todoID int) (common.Todo, error) {
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
			id = $1
			and user_id = $2;`, todoID, userID)
	if err != nil {
		return common.Todo{}, err
	}

	return todo, nil
}

func (tr *TodoRepository) GetAll(ctx context.Context, userID int, pagination common.Pagination) ([]common.Todo, error) {
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
			user_id = $1
		offset $2 limit $3;`, userID, pagination.CurrentPage, pagination.ItemsPerPage)
	if err != nil {
		return []common.Todo{}, err
	}

	return todos, nil
}

func (tr *TodoRepository) Update(ctx context.Context, todo common.Todo) error {
	queryString := fmt.Sprintf(`update todos set
									%s
								where
									id = :id`,
		postgresql.UpdateConditionFromStruct(&todo))

	res, err := tr.db.NamedExecContext(ctx, queryString, todo)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Rows not affected")
	}

	return nil
}

func (tr *TodoRepository) Delete(ctx context.Context, userID, todoID int) error {
	res, err := tr.db.ExecContext(ctx,
		`delete from todos where id = $1 and user_id = $2;`, todoID, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Rows not affected")
	}

	return nil
}

func (tr *TodoRepository) GetUserIDFromTodo(ctx context.Context, todoID int) (int, error) {
	row := tr.db.QueryRowContext(ctx,
		`select user_id from todos where id = $1`,
		todoID)

	var userID int
	if err := row.Scan(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}
