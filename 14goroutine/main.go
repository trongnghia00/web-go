package main

import (
	"fmt"
	"time"
)

func foo() {
	for i := 0; i <= 3; i++ {
		fmt.Println("foo: ", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func bar() {
	for i := 0; i <= 3; i++ {
		fmt.Println("bar: ", i)
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	go foo()
	go bar()

	time.Sleep(time.Millisecond * 900)
}
