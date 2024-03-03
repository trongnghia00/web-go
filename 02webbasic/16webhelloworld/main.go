package main

import (
	"fmt"
	"myapp/pkg/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application at port 8080 .....")
	http.ListenAndServe(":8080", nil)
}
