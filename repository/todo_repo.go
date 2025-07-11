package repository

import (
	"todolist/db"
	"todolist/models"
)

type TodoRepository interface {
	GetTodoByName(name string) (*models.ToDoList, error)
	CreateToDo(todo *models.ToDoList) error
	GetAllTodos() ([]models.ToDoList, error)
}

type TodoRepo struct{}

func (r *TodoRepo) GetTodoByName(name string) (*models.ToDoList, error) {
	var todo models.ToDoList
	err := db.DB.Where("name = ?", todo.Name).First(&todo).Error
	if err != nil {
		return &models.ToDoList{}, err
	}

	return &todo, nil

}

func (r *TodoRepo) CreateToDo(todo *models.ToDoList) error {
	err := db.DB.Create(&todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepo) GetAllTodos() ([]models.ToDoList, error) {
	var todos []models.ToDoList
	err := db.DB.Find(&todos).Error
	if err != nil {
		return []models.ToDoList{}, err
	}

	return todos, nil
}
