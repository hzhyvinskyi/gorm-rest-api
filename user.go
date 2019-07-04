package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
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
	db, err := gorm.Open("sqlite3", "appdb")
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to the database")
	}
	defer db.Close()

	var vars = mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})

	fmt.Fprintf(w, "User was created successfully")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "appdb")
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to the database")
	}
	defer db.Close()

	var vars = mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User

	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)

	fmt.Fprintf(w, "User was successfully updated")
}

func deletedUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "appdb")
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to the database")
	}
	defer db.Close()

	var vars = mux.Vars(r)
	name := vars["name"]

	var user User

	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "User was deleted")
}
