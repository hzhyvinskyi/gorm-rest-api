package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name string
	Email string
}

func initMigration() {
	db, err := gorm.Open("sqlite3", "appdb")
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to the database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

// Fetches all users from the DB and returns a list of them.
func getUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "appdb")
	if err != nil {
		panic("Cannot connect to the database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
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
