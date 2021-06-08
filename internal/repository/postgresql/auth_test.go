package postgesql

import (
	"context"
	"errors"
	"fmt"
	"reflect"
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
				rows := sqlmock.NewRows([]string{"id"}).AddRow("asd")
				mock.ExpectQuery("insert into users").WithArgs("name", "username", "email", "password").WillReturnRows(rows)
			},
			wantErr: true,
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
				rows := sqlmock.NewRows([]string{"id"}).AddRow(0)
				mock.ExpectQuery("insert into users").WithArgs("name", "username", "email", "password").WillReturnRows(rows)
			},
			want: 0,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			got, err := tc.r.CreateUser(context.Background(), tc.user)
			if (err != nil) != tc.wantErr {
				fmt.Println("test case name:", tc.name)
				t.Errorf("Get() error new = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if err == nil && got != tc.want {
				fmt.Println("test case name:", tc.name)
				t.Errorf("Get() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAuthorizationRepository_CreateUserSession(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	ar := NewAuthorizationRepository(sqlxDB)

	testTable := []struct {
		name            string
		r               *AuthorizationRepository
		userID          int
		userIP          string
		refreshToken    string
		refreshTokenTTL int64
		mock            func()
		want            int
		wantErr         bool
	}{
		{
			name:            "OK",
			r:               ar,
			userID:          1,
			userIP:          "IP",
			refreshToken:    "refreshToken",
			refreshTokenTTL: 1,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("insert into user_sessions").
					WithArgs(1, "IP", "refreshToken", 1).
					WillReturnRows(rows)
			},
			want: 1,
		},
		{
			name:            "Insert error",
			r:               ar,
			userID:          1,
			userIP:          "IP",
			refreshToken:    "refreshToken",
			refreshTokenTTL: 1,
			mock: func() {
				mock.ExpectQuery("insert into user_sessions").
					WithArgs(1, "IP", "refreshToken", 1).
					WillReturnError(errors.New("insert error"))
			},
			wantErr: true,
		},
		{
			name:            "User does not exist",
			r:               ar,
			userID:          -1,
			userIP:          "IP",
			refreshToken:    "refreshToken",
			refreshTokenTTL: 1,
			mock: func() {
				mock.ExpectQuery("insert into user_sessions").
					WithArgs(-1, "IP", "refreshToken", 1).
					WillReturnError(errors.New("user does not exist"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			got, err := tc.r.CreateUserSession(context.Background(), tc.userID, tc.userIP, tc.refreshToken, tc.refreshTokenTTL)
			if (err != nil) != tc.wantErr {
				fmt.Println("test case name:", tc.name)
				t.Errorf("Get() error new = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if err == nil && got != tc.want {
				fmt.Println("test case name:", tc.name)
				t.Errorf("Get() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAuthorizationRepository_UpdateUserSession(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	ar := NewAuthorizationRepository(sqlxDB)

	testTable := []struct {
		name            string
		r               *AuthorizationRepository
		userID          int
		refreshToken    string
		refreshTokenTTL int64
		mock            func()
		want            int
		wantErr         bool
	}{
		{
			name:            "OK",
			r:               ar,
			userID:          1,
			refreshToken:    "refreshToken",
			refreshTokenTTL: 1,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("update user_sessions set").
					WithArgs("refreshToken", 1, 1).
					WillReturnRows(rows)
			},
			want: 1,
		},
		{
			name:            "Insert error",
			r:               ar,
			userID:          1,
			refreshToken:    "refreshToken",
			refreshTokenTTL: 1,
			mock: func() {
				mock.ExpectQuery("update user_sessions set").
					WithArgs("refreshToken", 1, 1).
					WillReturnError(errors.New("update error"))
			},
			wantErr: true,
		},
		{
			name:            "User does not exist",
			r:               ar,
			userID:          -1,
			refreshToken:    "refreshToken",
			refreshTokenTTL: 1,
			mock: func() {
				mock.ExpectQuery("update user_sessions set").
					WithArgs("refreshToken", -1, 1).
					WillReturnError(errors.New("user does not exist"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			got, err := tc.r.UpdateUserSession(context.Background(), tc.userID, tc.refreshToken, tc.refreshTokenTTL)
			if (err != nil) != tc.wantErr {
				fmt.Println("test case name:", tc.name)
				t.Errorf("Get() error new = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if err == nil && got != tc.want {
				fmt.Println("test case name:", tc.name)
				t.Errorf("Get() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAuthorizationRepository_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	ar := NewAuthorizationRepository(sqlxDB)

	testTable := []struct {
		name     string
		r        *AuthorizationRepository
		username string
		password string
		mock     func()
		want     common.User
		wantErr  bool
	}{
		{
			name:     "OK",
			r:        ar,
			username: "username",
			password: "password",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "username", "email", "password"}).
					AddRow(1, "name", "username", "email", "password")
				mock.ExpectQuery("select id, name, username, email, password from users").
					WithArgs("username", "password").
					WillReturnRows(rows)
			},
			want: common.User{
				ID:       1,
				Name:     "name",
				Username: "username",
				Email:    "email",
				Password: "password",
			},
		},
		{
			name:     "select error",
			r:        ar,
			username: "username",
			password: "password",
			mock: func() {
				mock.ExpectQuery("select id, name, username, email, password from users").
					WithArgs("username", "password").WillReturnError(errors.New("select error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			got, err := tc.r.GetUser(context.Background(), tc.username, tc.password)
			if (err != nil) != tc.wantErr {
				fmt.Println("test case name:", tc.name)
				t.Errorf("Get() error new = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tc.want) {
				fmt.Println("test case name:", tc.name)
				t.Errorf("Get() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAuthorizationRepository_GetUserSession(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	ar := NewAuthorizationRepository(sqlxDB)

	testTable := []struct {
		name         string
		r            *AuthorizationRepository
		userID       int
		refreshToken string
		mock         func()
		want         common.UserSession
		wantErr      bool
	}{
		{
			name:         "OK",
			r:            ar,
			userID:       1,
			refreshToken: "refreshToken",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "user_ip", "refresh_token", "refresh_token_ttl"}).
					AddRow(1, 1, "IP", "refreshToken", 1)
				mock.ExpectQuery("select id, user_id, user_ip, refresh_token, refresh_token_ttl from user_session").
					WithArgs(1, "refreshToken").
					WillReturnRows(rows)
			},
			want: common.UserSession{
				ID:              1,
				UserID:          1,
				UserIP:          "IP",
				RefreshToken:    "refreshToken",
				RefreshTokenTTL: 1,
			},
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			got, err := tc.r.GetUserSession(context.Background(), tc.userID, tc.refreshToken)
			if (err != nil) != tc.wantErr {
				fmt.Println("test case name:", tc.name)
				t.Errorf("Get() error new = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tc.want) {
				fmt.Println("test case name:", tc.name)
				t.Errorf("Get() = %v, want %v", got, tc.want)
			}
		})
	}
}
