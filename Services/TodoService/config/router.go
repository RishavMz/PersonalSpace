package config

import (
	"TodoService/controllers"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func CreateRouter(dbConnection *gorm.DB) http.Handler {
	dbConn = dbConnection
    router := mux.NewRouter()

	router.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) { controllers.GetAllTodo(w, r, dbConn) }).Methods("GET")
	router.HandleFunc("/todoByUser", func(w http.ResponseWriter, r *http.Request) { controllers.GetTodoByUser(w, r, dbConn) }).Methods("GET")
	router.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) { controllers.CreateTodo(w, r, dbConn) }).Methods("POST")

	return router
}
