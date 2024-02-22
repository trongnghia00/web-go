package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.page.tmpl")
}

func renderTemplates(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := devideValues(100.0, 0.0)
// 	if err != nil {
// 		fmt.Fprintf(w, "Cannot divide by zero")
// 		return
// 	}
// 	fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 0.0, f)
// }

// func addValues(x, y int) int {
// 	return x + y
// }

// func devideValues(x, y float32) (float32, error) {
// 	if y == 0 {
// 		err := errors.New("cannot divide by zero")
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil
// }

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	// http.HandleFunc("/divide", Divide)

	fmt.Println("Starting application at port 8080 .....")
	http.ListenAndServe(":8080", nil)
}
