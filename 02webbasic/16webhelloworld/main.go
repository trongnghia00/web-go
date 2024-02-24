package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting application at port 8080 .....")
	http.ListenAndServe(":8080", nil)
}
