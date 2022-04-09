package service

import (
	todoappserver "github.com/VMironiuk/todo-app-server"
	"github.com/VMironiuk/todo-app-server/pkg/repository"
)

type TodoItemService struct {
	repository     repository.TodoItem
	listRepository repository.TodoList
}

func NewTodoItemService(repository repository.TodoItem, listRepository repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repository:     repository,
		listRepository: listRepository,
	}
}

func (s *TodoItemService) Create(userId, listId int, item todoappserver.TodoItem) (int, error) {
	_, err := s.listRepository.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repository.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todoappserver.TodoItem, error) {
	return s.repository.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (todoappserver.TodoItem, error) {
	return s.repository.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repository.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input todoappserver.UpdateItemInput) error {
	return s.repository.Update(userId, itemId, input)
}
