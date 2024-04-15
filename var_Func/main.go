package main

import "fmt"

func main() {
	//fmt.Println("Hello world...")

	//var whatToSay string
	//var i int

	//whatToSay = "Goodbye"
	//fmt.Println(whatToSay)

	//i = 23
	//fmt.Println("i is set to", i)

	whatAsSaid, moreSaid := saySomething()
	fmt.Println("The function said:", whatAsSaid, moreSaid)
}

func saySomething() (string, string) {
	return "Hi", "there"
}
