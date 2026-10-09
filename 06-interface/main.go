package main

import (
	"fmt"
)

func describeValue(value any) string {
	switch v := value.(type) {
	case string:
		return fmt.Sprintf("string: %v", v)
	case int:
		return fmt.Sprintf("int: %d", v)
	case bool:
		return fmt.Sprintf("bool: %t", v)
	default:
		return "unknown type"
	}
}

func main() {
	fmt.Println(describeValue("Go"))
	fmt.Println(describeValue(42))
	fmt.Println(describeValue(true))
	fmt.Println(describeValue(3.14))
}
