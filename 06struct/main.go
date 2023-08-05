package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

type rectangle struct {
	width  float64
	height float64
}

func (rec *rectangle) Area() float64 {
	return rec.width * rec.height
}

func main() {
	user := Person{
		Name: "Nghia",
		Age:  20,
	}

	fmt.Println(user.Name, user.Age)
	fmt.Println(user.Email)

	var rec1 rectangle
	rec1.width = 10
	rec1.height = 5

	fmt.Println("Area:", rec1.Area())
}
