package main

import "fmt"

type Box[T any] struct {
	Value T
}

type Number interface {
	~int | ~int64 | ~float64
}

type Score int

func Double[T Number](value T) T {
	return value * 2
}

func First[T any](values []T) (T, bool) {
	var zero T

	if len(values) == 0 {
		return zero, false
	}

	return values[0], true
}

func Contains[T comparable](values []T, target T) bool {
	for _, v := range values {
		if v == target {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(First([]int{10, 20, 30}))
	fmt.Println(First([]string{"Go", "Rust"}))
	fmt.Println(First([]int{}))

	fmt.Println(Contains([]int{1, 2, 3}, 2))
	fmt.Println(Contains([]int{1, 2, 3}, 4))
	fmt.Println(Contains([]string{"go", "rust"}, "go"))

	fmt.Println(Double(2))
	fmt.Println(Double(5.5))
	fmt.Println(Double[int64](5))

	fmt.Println(Double(Score(10)))

	intBox := Box[int]{Value: 42}
	stringBox := Box[string]{Value: "Go"}
	fmt.Printf("%d\n%s", intBox.Value, stringBox.Value)
}
