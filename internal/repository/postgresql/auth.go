package postgesql

import (
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

func (ar *AuthorizationRepository) CreateUser(user common.User) (int, error) {
	row := ar.db.QueryRow(`
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

func (ar *AuthorizationRepository) GetUser(username, password string) (common.User, error) {
	var user common.User
	err := ar.db.Get(&user, `
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
