package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/EraldCaka/server-go/api"
)

func main() {
	// server := api.NewApiServer(":3423")
	fmt.Println("hi!")
	// server.Start()
	database := NewDatabase()
	newJson, err := json.MarshalIndent(database, "", "     ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(newJson))
}

type Database struct {
	dbconn   string `json:"dbconn"`
	username string `json:"username"`
	password string `json:"password"`
}

type Storer interface {
	NewDatabase() *Database
}

func NewDatabase() *Database {
	return &Database{
		dbconn:   "conn",
		username: "username",
		password: "password",
	}
}
