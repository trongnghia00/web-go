package main

import (
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

func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting application at port 8080 .....")
	http.ListenAndServe(":8080", nil)
}
