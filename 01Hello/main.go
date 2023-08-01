package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var newLine string
	var i int

	newLine = "Goodbye world !"
	i = 10

	fmt.Println(newLine)
	fmt.Println("i is set to", i)

	someThing, someThingElse := saySomething()

	fmt.Println("The func returned", someThing, someThingElse)
}

func saySomething() (string, string) {
	return "something", "else"
}
