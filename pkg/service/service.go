package service

import (
	goRest "go-rest"
	"go-rest/pkg/repository"
)

type Authorization interface {
	CreateUser(user goRest.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
