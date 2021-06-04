package postgesql

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/asetriza/golang-web-app/internal/common"
	"github.com/jmoiron/sqlx"
)

func TestTodoRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	tr := NewTodoRepository(sqlxDB)

	testTable := []struct {
		name    string
		r       *TodoRepository
		user    common.Todo
		mock    func()
		want    int
		wantErr bool
	}{
		{
			name: "OK",
			r:    tr,
			user: common.Todo{
				UserID:      1,
				Name:        "name",
				Description: "description",
				NotifyDate:  1,
				Done:        false,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("insert into todos").WithArgs(1, "name", "description", 1, false).WillReturnRows(rows)
			},
			want: 1,
		},
		{
			name: "Empty fields",
			r:    tr,
			user: common.Todo{
				UserID:      1,
				Name:        "name",
				Description: "description",
				NotifyDate:  1,
				Done:        false,
			},
			mock: func() {
				mock.ExpectQuery("insert into todos").WithArgs(1, "name", "description", 1, false).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			got, err := tc.r.Create(context.Background(), tc.user)
			if (err != nil) != tc.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if err == nil && got != tc.want {
				t.Errorf("Get() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAuthorizationRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	tr := NewTodoRepository(sqlxDB)

	testTable := []struct {
		name       string
		r          *TodoRepository
		userID     int
		pagination common.Pagination
		mock       func()
		want       []common.Todo
		wantErr    bool
	}{
		{
			name:   "OK",
			r:      tr,
			userID: 1,
			pagination: common.Pagination{
				CurrentPage:  1,
				ItemsPerPage: 2,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "name", "description", "notify_date", "done"}).
					AddRow(1, 1, "name", "description", 1, false).
					AddRow(2, 1, "name", "description", 1, false)
				mock.ExpectQuery("select id, user_id, name, description, notify_date, done from todos").
					WithArgs(1, 0, 2).
					WillReturnRows(rows)
			},
			want: []common.Todo{
				{
					ID:          1,
					UserID:      1,
					Name:        "name",
					Description: "description",
					NotifyDate:  1,
					Done:        false,
				},
				{
					ID:          2,
					UserID:      1,
					Name:        "name",
					Description: "description",
					NotifyDate:  1,
					Done:        false,
				},
			},
		},
		{
			name:   "Incorrect user id",
			r:      tr,
			userID: -1,
			mock: func() {
				mock.ExpectQuery("select id, user_id, name, description, notify_date, done from todos").
					WithArgs(1).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			got, err := tc.r.GetAll(context.Background(), tc.userID, tc.pagination.CalculateOffset())
			if (err != nil) != tc.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Get() = %v, want %v", got, tc.want)
				return
			}
		})
	}
}

func TestAuthorizationRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	tr := NewTodoRepository(sqlxDB)

	testTable := []struct {
		name    string
		r       *TodoRepository
		userID  int
		todoID  int
		mock    func()
		want    common.Todo
		wantErr bool
	}{
		{
			name:   "OK",
			r:      tr,
			userID: 1,
			todoID: 1,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "name", "description", "notify_date", "done"}).
					AddRow(1, 1, "name", "description", 1, false)
				mock.ExpectQuery("select id, user_id, name, description, notify_date, done from todos").
					WithArgs(1, 1).
					WillReturnRows(rows)
			},
			want: common.Todo{
				ID:          1,
				UserID:      1,
				Name:        "name",
				Description: "description",
				NotifyDate:  1,
				Done:        false,
			},
		},
		{
			name:   "Empty fields",
			r:      tr,
			todoID: -1,
			userID: 1,
			mock: func() {
				mock.ExpectQuery("select id, user_id, name, description, notify_date, done from todos").
					WithArgs(-1, 1).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			got, err := tc.r.Get(context.Background(), tc.userID, tc.todoID)
			if (err != nil) != tc.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Get() = %v, want %v", got, tc.want)
				return
			}
		})
	}
}

func TestAuthorizationRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	tr := NewTodoRepository(sqlxDB)

	testTable := []struct {
		name    string
		r       *TodoRepository
		todoID  int
		mock    func()
		wantErr bool
	}{
		{
			name:   "OK",
			r:      tr,
			todoID: 1,
			mock: func() {
				rows := sqlmock.NewRows([]string{})
				mock.ExpectQuery("delete from todos").
					WithArgs(1).
					WillReturnRows(rows).
					RowsWillBeClosed()
			},
		},
		{
			name:   "Error no row",
			r:      tr,
			todoID: -1,
			mock: func() {
				mock.ExpectQuery("delete from todos").WithArgs(-1).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := tc.r.Delete(context.Background(), tc.todoID)
			if (err != nil) != tc.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
