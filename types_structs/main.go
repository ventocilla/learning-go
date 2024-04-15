package main

import (
	// "log"
	"fmt"
	"time"
)

//var s = "seven"

type User struct {
	Firstname   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {

	user := User{
		Firstname: "LV",
		LastName:  "Vento",
	}

	fmt.Println(user.LastName)
	fmt.Println(user.BirthDate)

	/*
		var s2 = "six"
		s := "eight"
		log.Println("s is", s)
		log.Println("s2 is", s2)
		saySomething("xxx")
	*/
}

/*
func saySomething(s3 string) (string, string) {
	log.Println("s3 inside func", s)
	return s3, "World"
}
*/
