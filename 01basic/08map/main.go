package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	studentAge := make(map[string]int)

	studentAge["Nghia"] = 22
	studentAge["Trong"] = 25

	fmt.Println(studentAge["Nghia"], studentAge["Trong"])

	people := make(map[string]Person)
	p1 := Person{Name: "Nghia", Age: 22}
	people["Nghia"] = p1
	fmt.Println("Name:", people["Nghia"].Name, "Age:", people["Nghia"].Age)

	ages := make(map[int]string)
	ages[22] = "Nghia"
	fmt.Println("Age: 22, Name:", ages[22])
}
