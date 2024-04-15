package main

import "fmt"

func main() {
	letter := "C"

	switch letter {
	case "A":
	case "C":
		fmt.Println("Letter A or C")
	case "B":
		fmt.Println("Letter B")
	default:
		fmt.Println("Another letter")
	}
	fmt.Println("End of File")
}
