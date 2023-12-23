package config

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"TodoService/models"
)

var users []models.User

type ResponseData struct {
	Message string `json:"message"`
}

func CreateRouter() http.Handler {
    router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	return router
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting a list of users")
	w.Header().Set("Content-Type", "application/json")
	responseData := ResponseData{
		Message: "Hello, World!",
	}

	// Convert the response data to JSON
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		// Handle the error if JSON marshaling fails
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the client
	w.Write(jsonData)
}
