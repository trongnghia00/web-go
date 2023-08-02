package main

import "fmt"

func main() {
	var myColor string
	myColor = "Green"

	fmt.Println("My color is:", myColor)
	changeUsingPointer(&myColor)
	fmt.Println("after func call, my color is", myColor)
}

func changeUsingPointer(s *string) {
	fmt.Println("s:", *s)
	newColor := "Red"
	*s = newColor
}
