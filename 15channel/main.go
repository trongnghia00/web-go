package main

import (
	"fmt"
	"sync"
)

func sum(numbers []int, ch chan int) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	ch <- total
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resultChannel := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		sum(numbers[:len(numbers)/2], resultChannel)
	}()

	go func() {
		defer wg.Done()
		sum(numbers[len(numbers)/2:], resultChannel)
	}()

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	finalSum := 0
	for sum := range resultChannel {
		finalSum += sum
	}

	fmt.Println("Sum is: ", finalSum)
}
