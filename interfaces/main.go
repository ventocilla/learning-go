package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Bread string
}

type Gorila struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Bread: "German Shephere",
	}
	PrintInfo(&dog)

	gorilla := Gorila{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorila) Says() string {
	return "Ugh"
}

func (g *Gorila) NumberOfLegs() int {
	return 2
}
