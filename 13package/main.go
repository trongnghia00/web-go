package main

import (
	"fmt"
	"mymodule/myutils"
)

func main() {
	result := myutils.Factorial(5)
	fmt.Println("5!=", result)
}
