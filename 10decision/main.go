package main

import "fmt"

func main() {
	age := 5
	if age >= 18 {
		fmt.Println("Bạn đã đủ tuổi lái xe.")
	} else {
		fmt.Println("Bạn chưa đủ tuổi")
	}

	score := 65
	if score >= 85 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else if score >= 55 {
		fmt.Println("C")
	} else if score >= 40 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

	age = 25
	isStudent := true

	if age >= 18 && age <= 30 || !isStudent {
		// fmt.Println("Bạn là thanh niên và là sinh viên")
	}

	day := "Sunday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Ngày làm việc")
	case "Saturday", "Sunday":
		fmt.Println("Ngày cuối tuần")
	default:
		fmt.Println("Không xác định")
	}
}
