package config

import (
	"TodoService/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func CreateRouter(dbConnection *gorm.DB) http.Handler {
	dbConn = dbConnection
    router := mux.NewRouter()

	router.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) { handlers.GetAllTodo(w, r, dbConn) }).Methods("GET")
	router.HandleFunc("/todoByUser", func(w http.ResponseWriter, r *http.Request) { handlers.GetTodoByUser(w, r, dbConn) }).Methods("GET")
	router.HandleFunc("/todoUpdate", func(2 http.ResponseWriter, r *http.Request) { handlers.UpdateTodo(w, r, dbConn) }).Methods("PUT")
	router.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) { handlers.CreateTodo(w, r, dbConn) }).Methods("POST")

	return router
}
