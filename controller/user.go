package controller

import (
	"fmt"
	database "github.com/fooksupachai/Restful-Golang-Mongo/database"
	"net/http"
)

// GetUser function
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		result := database.GetMember()
		fmt.Fprint(w, result)
	case "POST":
		w.Write([]byte(`{"message": [{"createUser": "1"}]}`))
	}
}
