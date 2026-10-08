package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		//return учебный проект не завершаем работу
	} else {
		fmt.Printf("10/2: %v\n", result)
	}

	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
		//return
	} else {
		fmt.Printf("10/0: %v\n", result2)
	}

	min, max, err := minMax([]int{10, 20, 30, 40})
	if err != nil {
		fmt.Println("Error:", err)
		//return
	} else {
		fmt.Printf("Min: %v, Max: %v\n", min, max)
	}

	min2, max2, err2 := minMax([]int{})
	if err2 != nil {
		fmt.Println("Error:", err2)
		//return
	} else {
		fmt.Printf("Min: %v, Max: %v\n", min2, max2)
	}

	min3, max3, err3 := minMax([]int{30, 10, 40, 20})
	if err3 != nil {
		fmt.Println("Error:", err3)
		//return
	} else {
		fmt.Printf("Min: %v, Max: %v\n", min3, max3)
	}

	for _, age := range []int{-5, 25, 200} {
		err := validateAge(age)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("Age is valid")
	}
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func validateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}

	if age > 150 {
		return errors.New("age is too large")
	}

	return nil
}

func minMax(numbers []int) (int, int, error) {
	if len(numbers) == 0 {
		return 0, 0, errors.New("slice is empty")
	}

	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}
	return min, max, nil
}
