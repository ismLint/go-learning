package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
}

type Contact struct {
	Name  string
	Phone string
}

var contactMap map[string]Contact

func main() {
	words := []string{"go", "rust", "go", "python", "go", "rust"}
	countWords := wordFrequency(words)
	fmt.Println(countWords)

	users := []User{
		{Name: "Nick", Email: "nick@example.com", Age: 25},
		{Name: "Jason", Email: "jason@example.com", Age: 17},
		{Name: "John", Email: "john@example.com", Age: 30},
		{Name: "", Email: "", Age: 0}, //пустой срез для теста
	}
	adultUsers := findAdultUsers(users)
	fmt.Println(adultUsers)

	contactMap = map[string]Contact{
		"Nick":  {Name: "Nick", Phone: "123-456-7890"},
		"Jason": {Name: "Jason", Phone: "987-654-3210"},
		"John":  {Name: "John", Phone: ""},
	}

	findContact1, ok1 := findContact(contactMap, "Nick")
	if ok1 {
		fmt.Println(findContact1)
	} else {
		fmt.Println("Contact not found")
	}

	findContact2, ok2 := findContact(contactMap, "John")
	if ok2 {
		fmt.Println(findContact2)
	} else {
		fmt.Println("Contact not found")
	}

	contact := Contact{Name: "John", Phone: ""}
	contact.Phone = "123-123-1234"
	contactMap["John"] = contact

	findContact3, ok3 := findContact(contactMap, "John")
	if ok3 {
		fmt.Println(findContact3)
	} else {
		fmt.Println("Contact not found")
	}

	delete(contactMap, "Jason")
	findContact4, ok4 := findContact(contactMap, "Jason")
	if ok4 {
		fmt.Println(findContact4)
	} else {
		fmt.Println("Contact not found")
	}
}

func wordFrequency(words []string) map[string]int {
	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}

	return frequency
}

func findAdultUsers(users []User) []User {
	var adultUsers []User
	for _, user := range users {
		if user.Age >= 18 {
			adultUsers = append(adultUsers, user)
		}
	}
	return adultUsers
}

func findContact(contacts map[string]Contact, name string) (Contact, bool) {
	contact, ok := contacts[name]

	if ok {
		return contact, true
	}

	return Contact{}, false
}
