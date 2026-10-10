package main

import (
	"fmt"

	"go-learning/packages/user"
)

func main() {
	currentUser := []user.User{
		user.NewUser("Nick", 25),
		user.NewUser("James", 17),
		user.NewUser("Vadim", 18),
	}
	fmt.Println(currentUser)

	fmt.Println(user.IsAdult(17))
	fmt.Println(user.IsAdult(18))
	fmt.Println(user.IsAdult(25))

	fmt.Println(currentUser[0].IsAdult())

	fmt.Println(user.FilterAdults(currentUser))
}
