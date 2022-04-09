package repository

import (
	todoappserver "github.com/VMironiuk/todo-app-server"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todoappserver.User) (int, error)
	GetUser(username, password string) (todoappserver.User, error)
}

type TodoList interface {
	Create(userId int, list todoappserver.TodoList) (int, error)
	GetAll(userId int) ([]todoappserver.TodoList, error)
	GetById(userId, listId int) (todoappserver.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todoappserver.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item todoappserver.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todoappserver.TodoItem, error)
	GetById(userId, itemId int) (todoappserver.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todoappserver.UpdateItemInput) error
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
		TodoItem:      NewTodoItemPostgres(db),
	}
}
