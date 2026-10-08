package main

import (
	"errors"
	"fmt"
)

func main() {
	numbers := []int{10, 20, 30, 40, 5, 25, 35}
	numbersWithoutEven := []int{1, 3, 5, 7, 9}
	numbersEmpty := []int{}

	result := double(numbers)
	fmt.Printf("Doubled numbers are: %v\n", result)

	result2 := filterEven(numbers)
	fmt.Printf("Filtered even numbers are: %v\n", result2)

	result3 := filterEven(numbersWithoutEven)
	fmt.Printf("Filtered even numbers are: %v\n", result3)

	result4 := filterEven(numbersEmpty)
	fmt.Printf("Filtered even numbers are: %v\n", result4)

	result5, err5 := removeAt(numbers, 2)
	if err5 != nil {
		fmt.Println("Error: ", err5)
		return
	}
	fmt.Printf("Slice after removing element at index 2: %v\n", result5)

	result6, err6 := removeAt(numbers, 0)
	if err6 != nil {
		fmt.Println("Error: ", err6)
		return
	}
	fmt.Println("Slice after removing element at index 0: ", result6)

	result7, err7 := removeAt(numbers, len(numbers)-1)
	if err7 != nil {
		fmt.Println("Error: ", err7)
		return
	}
	fmt.Println("Slice after removing element at index len(numbers)-1: ", result7)

	result8, err8 := removeAt(numbers, len(numbers))
	if err8 != nil {
		fmt.Println("Error: ", err8)
		//return
	}
	fmt.Println("Slice after removing element at index len(numbers): ", result8)

	result9, err9 := removeAt(numbers, 2)
	if err9 != nil {
		fmt.Println("Error: ", err9)
		//return учебный проект пропускаем ошибку
	}
	fmt.Println("Original:", numbers)
	fmt.Println("Result: ", result9)

	result10, err10 := removeAt(numbers, -1)
	if err10 != nil {
		fmt.Println("Error: ", err10)
		//return
	}
	fmt.Println("Result: ", result10)
}

func double(numbers []int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = num * 2
	}
	return result
}

func filterEven(numbers []int) []int {
	result := make([]int, 0, len(numbers))
	for _, num := range numbers {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result
}

func removeAt(numbers []int, index int) ([]int, error) {
	if index < 0 || index >= len(numbers) {
		return numbers, errors.New("index out of range")
	}

	result := make([]int, 0, len(numbers)-1)
	result = append(result, numbers[:index]...)
	result = append(result, numbers[index+1:]...)

	return result, nil
}
