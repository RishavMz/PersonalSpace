package controllers

import (
	"TodoService/models"

	"encoding/json"

	"gorm.io/gorm"
)

func FetchAllTodos(dbConn *gorm.DB) ([]byte, error) {
	var todos []models.Todo
	if err := dbConn.Find(&todos).Error; err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(todos)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func FetchTodosByUser(dbConn *gorm.DB, userID uint64) ([]byte, error) {
	var todos []models.Todo
	if err := dbConn.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(todos)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func CreateTodo(dbConn *gorm.DB, newTodo models.Todo) (uint, error) {
    if err := dbConn.Create(&newTodo).Error; err != nil {
        return 0, err
    }
    return newTodo.ID, nil
}

func UpdateTodo(dbConn *gorm.DB, todoID uint64, updatedTodo models.Todo) error {
	var existingTodo models.Todo
	if err := dbConn.First(&existingTodo, todoID).Error; err != nil {
		return err
	}

	existingTodo.Title = updatedTodo.Title
	existingTodo.Description = updatedTodo.Description
	existingTodo.Priority = updatedTodo.Priority
	existingTodo.Status = updatedTodo.Status

	if err := dbConn.Save(&existingTodo).Error; err != nil {
		return err
	}
	return nil
}
