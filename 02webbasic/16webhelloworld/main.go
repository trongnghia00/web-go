package main

import (
	"errors"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Word !!")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 3)
	fmt.Fprintf(w, "This is About page. 2 + 3 = %d", sum)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := devideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}
	fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 0.0, f)
}

func addValues(x, y int) int {
	return x + y
}

func devideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println("Starting application at port 8080 .....")
	http.ListenAndServe(":8080", nil)
}
