package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func createUser() {

	var email string
	var users []User

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	content := readData("./data/Users.json")

	json.Unmarshal(content, &users)

	for _, user := range users {

		if user.Email == email {
			log.Fatal("User already exists in the db")
		}

	}

	var password string

	fmt.Print("Enter your password")
	fmt.Scan(&password)

	user := User{email, password}
	users = append(users, user)

	content, _ = json.MarshalIndent(&users, "", "")

	writeData("./data/Users.json", content)

	println("Created the user")
}
