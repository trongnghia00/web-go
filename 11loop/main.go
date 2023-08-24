package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}

	// for _, animal := range animals {
	// 	fmt.Println(animal)
	// }

	// animals := make(map[string]string)
	// animals["dog"] = "Basu"
	// animals["cat"] = "Bong gon"
	// for animalType, animal := range animals {
	// 	fmt.Println(animalType, animal)
	// }

	// var line = "Hello !!!"
	// for i, c := range line {
	// 	fmt.Println(i, ":", c)
	// }

	type User struct {
		Name string
		Age  int
	}

	var users []User
	users = append(users, User{"Nghia", 20})
	users = append(users, User{"Trong", 25})
	users = append(users, User{"Nguyen", 30})

	for _, user := range users {
		fmt.Println(user.Name, user.Age)
	}
}
