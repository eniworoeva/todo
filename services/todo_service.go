package services

import (
	"todolist/models"
	"todolist/repository"

	"github.com/google/uuid"
)

type TodoService struct {
	Repo repository.TodoRepository
}

func (s *TodoService) CreateToDo(req *models.ToDoList) error {
	//check it doesn't exist in the db already
	_, err := s.Repo.GetTodoByName(req.Name)
	if err == nil {
		return err
	}

	myuuid := uuid.NewString()
	req.ID = myuuid

	// add to db using encoder
	err = s.Repo.CreateToDo(req)
	if err != nil {
		return err
	}

	return nil
}

func (s *TodoService) GetToDos() ([]models.ToDoList, error) {
	//retrieve todolist of that user
	todos, err := s.Repo.GetAllTodos()
	if err != nil {
		return nil, err
	}

	return todos, nil

}
