package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	database "github.com/fooksupachai/Restful-Golang-Mongo/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Account struct
type Account struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Member struct {
	Name string `json:"name"`
}

// GetAllUser function
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	c := database.InitialDB()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	heroes := database.GetAllMember(c, bson.M{})

	var accounts []Account
	for _, hero := range heroes {
		var accountList = Account{"Account", hero.A}
		accounts = append(accounts, accountList)
	}

	resp := struct {
		Data   []Account `json:"data"`
		Status int       `json:"status"`
	}{
		Data:   accounts,
		Status: http.StatusOK,
	}

	json.NewEncoder(w).Encode(resp)
}

// GetUser function
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := struct {
		Age  string `json:"age"`
		Name string `json:"name"`
	}{
		Age:  r.FormValue("age"),
		Name: r.FormValue("name"),
	}
	json.NewEncoder(w).Encode(resp)
}

// CreateUser function
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Use Method POST instead")

	case "POST":

		var mem Member
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&mem)

		if err == nil {

			w.WriteHeader(http.StatusOK)

			resp := struct {
				Name string `json:"name"`
			}{
				Name: mem.Name,
			}
			json.NewEncoder(w).Encode(resp)
		}

		if err != nil {

			w.WriteHeader(http.StatusBadRequest)

			resp := struct {
				Status int    `json:"status"`
				Notify string `json:"notify"`
			}{
				Status: http.StatusBadRequest,
				Notify: "Invalid type object",
			}

			json.NewEncoder(w).Encode(resp)
		}

	}
}
