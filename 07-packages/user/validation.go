package user

import "errors"

func IsAdult(age int) bool {
	return age >= 18
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func ValidateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}

	if age > 150 {
		return errors.New("age is too large")
	}
	return nil
}
