package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Type string
}

type Gorilla struct {
	Name  string
	Color string
}

func main() {
	dog := Dog{
		Name: "Basu",
		Type: "German Shephered",
	}

	gorilla := Gorilla{
		Name:  "Jack",
		Color: "grey",
	}

	PrintInfo(&dog)
	PrintInfo(gorilla)
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

func (g Gorilla) Says() string {
	return "Gruuu"
}

func (g Gorilla) NumberOfLegs() int {
	return 2
}
