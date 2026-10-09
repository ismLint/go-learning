package user

func IsAdult(age int) bool {
	return age >= 18
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}
