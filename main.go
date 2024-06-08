package main

import (
	"log"
	"net/http"
)



func main() {

	file := ConfigureLogger()

	defer file.Close()
	log.SetOutput(file)

	ReadEnvironmentVariables()

	http.HandleFunc("/", HelloWorld)

	RunServer()
}
