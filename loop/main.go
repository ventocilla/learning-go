package main

import "fmt"

func main() {
	// var firstLine = "Once upon a time"
	// firstLine = "y"

	type User struct {
		FirstName string
		LastName  string
		// Email     string
		// Age       string
	}

	var users []User

	users = append(users, User{FirstName: "Leo", LastName: "Vento"})
	users = append(users, User{FirstName: "Yanira", LastName: "Veronica"})
	users = append(users, User{FirstName: "Marly", LastName: "Dantas"})
	users = append(users, User{FirstName: "Malu", LastName: "Ventocilla"})

	for _, v := range users {
		fmt.Println(v.FirstName, v.LastName)
	}
}
