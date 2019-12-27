package router

import (
	controller "github.com/fooksupachai/Restful-Golang-Mongo/controller"
	"net/http"
)

func init() {
	http.HandleFunc("/user", controller.GetUser)
	http.HandleFunc("/users", controller.GetAllUser)
	http.HandleFunc("/createUser", controller.CreateUser)
}
