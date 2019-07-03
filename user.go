package main

import (
	"fmt"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of users...")
}

func addUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New user...")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Updated user...")
}

func deletedUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleted user...")
}
