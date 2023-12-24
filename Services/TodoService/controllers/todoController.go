package controllers

import (
	"TodoService/handlers"
	"TodoService/models"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func GetAllTodo(w http.ResponseWriter, r *http.Request, dbConn *gorm.DB) {
	todosJSON, err := handlers.FetchAllTodos(dbConn)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(todosJSON)
}

func GetTodoByUser(w http.ResponseWriter, r *http.Request, dbConn *gorm.DB) {
	userIDParam := r.URL.Query().Get("userId")
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid userID parameter", http.StatusBadRequest)
		return
	}

	todosJSON, err := handlers.FetchTodosByUser(dbConn, userID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(todosJSON)
}

func CreateTodo(w http.ResponseWriter, r *http.Request, dbConn *gorm.DB) {
    var newTodo models.Todo
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&newTodo); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    id, err := handlers.CreateTodo(dbConn, newTodo)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to create todo: %v", err), http.StatusInternalServerError)
        return
    }

    newTodo.ID = id

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(newTodo); err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}

