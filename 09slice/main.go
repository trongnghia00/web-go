package main

import (
	"fmt"
	"sort"
)

func main() {
	var mySlice []string

	mySlice = append(mySlice, "Nghia")
	mySlice = append(mySlice, "Trong")

	fmt.Println(mySlice)

	var mySlice2 []int

	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 1)
	mySlice2 = append(mySlice2, 3)

	sort.Ints(mySlice2)

	fmt.Println(mySlice2)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(numbers[6:9])
}
