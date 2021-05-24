package postgesql

import (
	"context"
	"golang-web-app/internal/common"

	"github.com/jmoiron/sqlx"
)

type AuthorizationRepository struct {
	db *sqlx.DB
}

func NewAuthorizationRepository(db *sqlx.DB) *AuthorizationRepository {
	return &AuthorizationRepository{
		db: db,
	}
}

func (ar *AuthorizationRepository) CreateUser(ctx context.Context, user common.User) (int, error) {
	row := ar.db.QueryRowContext(ctx, `
		insert into users
			(name, username, password)
		values
			($1, $2, $3)
		returning
			id;
		`, user.Name, user.Username, user.Password)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ar *AuthorizationRepository) GetUser(ctx context.Context, username, password string) (common.User, error) {
	var user common.User
	err := ar.db.GetContext(ctx, &user, `
		select
			id
		from
			users
		where
			username = $1
			and password = $2;
		`, username, password)

	return user, err
}

func (ar *AuthorizationRepository) GetUserSession(ctx context.Context, userID int, refreshToken string) (common.UserSession, error) {
	var userSession common.UserSession
	err := ar.db.GetContext(ctx, &userSession, `
		select
			id,
			user_id,
			user_ip,
			refresh_token,
			refresh_token_ttl
		from
			user_sessions
		where
			user_id = $1
			and refresh_token = $2;
		`, userID, refreshToken)

	return userSession, err
}

func (ar *AuthorizationRepository) CreateUserSession(ctx context.Context, userID int, userIP, refreshToken string, refreshTokenTTL int64) (int, error) {
	row := ar.db.QueryRowContext(ctx, `
		insert into user_sessions
			(user_id, user_ip, refresh_token, refresh_token_ttl)
		values
			($1, $2, $3, $4)
		returning
			id;
		`, userID, userIP, refreshToken, refreshTokenTTL)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ar *AuthorizationRepository) UpdateUserSession(ctx context.Context, userID int, refreshToken string, refreshTokenTTL int64) (int, error) {
	row := ar.db.QueryRowContext(ctx, `
		update user_sessions set
			refresh_token = $1,
			refresh_token_ttl = $2
		where
			user_id = $3
		returning
			id;
		`, refreshToken, refreshTokenTTL, userID)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
