package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	go_rest "go-rest"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user go_rest.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT into %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (go_rest.User, error) {
	var user go_rest.User
	query := fmt.Sprintf("SELECT id from %s WHERE username=$1 AND password_hash=$2", userTable)

	err := r.db.Get(&user, query, username, password)

	return user, err
}
