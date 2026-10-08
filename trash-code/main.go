package main

import "fmt"

func main() {
	var name = "ismLint"
	var age = "19"
	var a = 5
	var b = 10
	hello(name, age)
	fmt.Println("Sum of", a, "and", b, "is", summ(a, b))
	var day = 3

	if a < b {
		fmt.Println("All ok!")
	} else {
		fmt.Println("All not ok!")
	}

	count := 10
	for count < 15 {
		count++
		fmt.Println("Count is", count)
	}

	switchDay(day)

	numbers := [3]int{10, 20, 30}
	fmt.Println("Number with 0 index is: ", numbers[0])

	numbers2 := []int{10, 20, 30}
	numbers2 = append(numbers2, 40)
	fmt.Println("Numbers are: ", numbers2)

	for index, value := range numbers2 {
		fmt.Println("Index is", index, "Value is", value)
	}

	for _, value := range numbers2 {
		fmt.Println("Value is", value)
	}
}

func hello(name string, age string) {
	fmt.Println("Hello", name, "Your age is", age)
}

func summ(a, b int) int {
	return a + b
}

func switchDay(day int) {
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	default:
		fmt.Println("Other day")
	}
}
