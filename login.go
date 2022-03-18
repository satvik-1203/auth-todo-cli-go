package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func login() {
	var email string
	var password string

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter your password: ")
	fmt.Scan(&password)

	content := readData("./data/Users.json")

	var users []User

	json.Unmarshal(content, &users)

	for _, user := range users {

		if user.Email != email {
			continue
		}

		if user.Password != password {
			println("Invalid password or usernamen")
		}

		payload := jwt.MapClaims{}
		payload["authorized"] = true
		payload["user_id"] = email
		payload["exp"] = time.Now().Add(time.Minute * 15).Unix()
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
		token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

		if err != nil {
			println("Something wen't wrong.")
			return
		}

		println("Your token is: ", token)
		return
	}

	fmt.Println("User doesn't exist in the database")
}
