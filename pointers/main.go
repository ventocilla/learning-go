package main

import (
	"fmt"
)

func main() {
	var myString = "Green"
	fmt.Println("MyString is set to:", myString)

	changeUsingPointer(&myString) // & to get reference to a pointer ..
	fmt.Println("After func call MyString is set to:", myString)
}

func changeUsingPointer(s *string) {
	fmt.Println("s is set to:", s)
	newValue := "Red"
	*s = newValue // * to refer to a pointer, to mutate it's value
}
