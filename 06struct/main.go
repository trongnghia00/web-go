package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

var number int

// private public
func main() {
	user := Person{
		Name: "Nghia",
		Age:  20,
	}

	fmt.Println(user.Name, user.Age)
	fmt.Println(user.Email)
}
