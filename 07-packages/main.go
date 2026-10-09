package main

import (
	"fmt"

	"go-learning/packages/user"
)

func main() {
	currentUser := user.NewUser("Nick", 25)
	fmt.Println(currentUser)

	fmt.Println(user.IsAdult(17))
	fmt.Println(user.IsAdult(18))
	fmt.Println(user.IsAdult(25))

	fmt.Println(currentUser.IsAdult())
}
