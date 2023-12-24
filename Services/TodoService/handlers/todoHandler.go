package handlers

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
