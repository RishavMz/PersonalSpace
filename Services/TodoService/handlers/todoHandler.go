package handlers

import (
	"TodoService/controllers"
	"TodoService/models"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllTodo(w http.ResponseWriter, r *http.Request, dbConn *gorm.DB) {
	todosJSON, err := controllers.FetchAllTodos(dbConn)
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

	todosJSON, err := controllers.FetchTodosByUser(dbConn, userID)
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

    id, err := controllers.CreateTodo(dbConn, newTodo)
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

func UpdateTodo(w http.ResponseWriter, r *http.Request, dbConn *gorm.DB) {
	vars := mux.Vars(r)
	todoID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}
	var updatedTodo models.Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedTodo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := controllers.UpdateTodo(dbConn, todoID, updatedTodo); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedTodo); err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}

