package main

import (
	"log"
	"net/http"

	_ "github.com/fooksupachai/Restful-Golang-Mongo/router"
)

func main() {
	log.Print("Rest server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
