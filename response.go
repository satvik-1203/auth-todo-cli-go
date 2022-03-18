package main

import "fmt"

func response() int {
	var option int

	fmt.Print("Enter your option: ")
	fmt.Scan(&option)

	fmt.Println()

	switch option {
	case 1:
		fmt.Println("Creating a user...")
		createUser()
	case 2:
		fmt.Println("Loggin in...")
		login()
	case 3:
		fmt.Println("Create a todo")
	case 4:
		fmt.Println("View todo")

	}
	return option
}
