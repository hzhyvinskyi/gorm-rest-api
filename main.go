package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", home).Methods("GET")
	myRouter.HandleFunc("/users", getUsers).Methods("GET")
	myRouter.HandleFunc("/users/{name}/{email}", addUser).Methods("POST")
	myRouter.HandleFunc("/users/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/users/{name}", deletedUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9091", myRouter))
}

func main() {
	fmt.Println("GORM REST API")

	initMigration()

	handleRequests()
}
