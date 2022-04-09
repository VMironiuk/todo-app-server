package service

import (
	todoappserver "github.com/VMironiuk/todo-app-server"
	"github.com/VMironiuk/todo-app-server/pkg/repository"
)

type Authorization interface {
	CreateUser(user todoappserver.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todoappserver.TodoList) (int, error)
	GetAll(userId int) ([]todoappserver.TodoList, error)
	GetById(userId, listId int) (todoappserver.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todoappserver.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item todoappserver.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todoappserver.TodoItem, error)
	GetById(userId, itemId int) (todoappserver.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todoappserver.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		TodoList:      NewTodoListService(repository.TodoList),
		TodoItem:      NewTodoItemService(repository.TodoItem, repository.TodoList),
	}
}
