package main

import (
	"TodoService/config"
	"net/http"

	"gorm.io/gorm"
)

var conn *gorm.DB
var router http.Handler

func main() {

	config.LoadEnvVariables()
	conn = config.CreateConnection()
	router = config.CreateRouter()
	config.StartServer(router)
}
