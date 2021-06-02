package postgesql

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/asetriza/golang-web-app/internal/common"
	"github.com/jmoiron/sqlx"
)

func TestAuthorizationRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	ar := NewAuthorizationRepository(sqlxDB)

	testTable := []struct {
		name    string
		r       *AuthorizationRepository
		user    common.User
		mock    func()
		want    int
		wantErr bool
	}{
		{
			name: "OK",
			r:    ar,
			user: common.User{
				Name:     "name",
				Username: "username",
				Email:    "email",
				Password: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("insert into users").WithArgs("name", "username", "email", "password").WillReturnRows(rows)
			},
			want: 1,
		},
		{
			name: "Empty fields",
			r:    ar,
			user: common.User{
				Name:     "",
				Username: "",
				Email:    "email",
				Password: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("insert into users").WithArgs("name", "username", "email", "password").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			got, err := tc.r.CreateUser(context.Background(), tc.user)
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
