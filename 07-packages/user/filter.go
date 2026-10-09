package user

func FilterAdults(users []User) []User {
	adultUser := make([]User, 0)

	for _, currentUser := range users {
		if currentUser.IsAdult() {
			adultUser = append(adultUser, currentUser)
		}
	}

	return adultUser
}
