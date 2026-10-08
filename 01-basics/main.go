package main

import "fmt"

func main() {
	name := "ismLint"
	age := 19
	describe(name, age)
	result1 := classifyNumber(1)
	result2 := classifyNumber(0)
	result3 := classifyNumber(-1)
	fmt.Printf("Result 1 is: %s\nResult 2 is: %s\nResult 3 is: %s\n", result1, result2, result3)

	enumerateNumbers()
}

func describe(name string, age int) {
	fmt.Printf("My name is %s and my age is %d\n", name, age)
}

func classifyNumber(number int) string {
	if number > 0 {
		return "positive"
	}

	if number == 0 {
		return "zero"
	}

	return "negative"
}

func enumerateNumbers() {
	for count := 1; count <= 20; count++ {
		if count%3 == 0 && count%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if count%3 == 0 {
			fmt.Println("Fizz")
		} else if count%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(count)
		}
	}
}
