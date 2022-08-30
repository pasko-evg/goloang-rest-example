package repository

import (
	"github.com/jmoiron/sqlx"
	goRest "go-rest"
)

type Authorization interface {
	CreateUser(user goRest.User) (int, error)
	GetUser(username, password string) (goRest.User, error)
}

type TodoList interface {
	Create(userId int, list goRest.TodoList) (int, error)
	GetAll(userId int) ([]goRest.TodoList, error)
	GetById(userId, listId int) (goRest.TodoList, error)
	Update(userId, listId int, input goRest.UpdateListInput) error
	Delete(userId, listId int) error
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
