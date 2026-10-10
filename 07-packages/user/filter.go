package user

func FilterAdults(users []User) []User {
	adults := make([]User, 0)

	for _, currentUser := range users {
		if currentUser.IsAdult() {
			adults = append(adults, currentUser)
		}
	}

	return adults
}
