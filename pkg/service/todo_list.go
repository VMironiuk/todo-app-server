package service

import (
	todoappserver "github.com/VMironiuk/todo-app-server"
	"github.com/VMironiuk/todo-app-server/pkg/repository"
)

type TodoListService struct {
	repository repository.TodoList
}

func NewTodoListService(repository repository.TodoList) *TodoListService {
	return &TodoListService{repository: repository}
}

func (s *TodoListService) Create(userId int, list todoappserver.TodoList) (int, error) {
	return s.repository.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todoappserver.TodoList, error) {
	return s.repository.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (todoappserver.TodoList, error) {
	return s.repository.GetById(userId, listId)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.repository.Delete(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input todoappserver.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repository.Update(userId, listId, input)
}
