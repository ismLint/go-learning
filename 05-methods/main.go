package main

import (
	"errors"
	"fmt"
)

type Counter struct {
	Value int
}

type User struct {
	Name string
	Age  int
}

type BankAccount struct {
	Owner   string
	Balance float64
}

func main() {
	number := 10
	increment(&number)
	fmt.Println("Incremented value:", number)

	users := []User{
		{Name: "Nick", Age: 19},
		{Name: "Vadim", Age: 18},
		{Name: "Nastya", Age: 18},
	}

	for index := range users {
		users[index].Birthday()
	}
	fmt.Println(users)

	counter := Counter{Value: 10}
	counter.Increment()

	fmt.Println(counter.Value)

	counter.IncrementCopy()
	fmt.Println(counter.Value) //11

	counter.Add(5)
	fmt.Println(counter.Value) // 16

	accounts := []BankAccount{
		{Owner: "Nick", Balance: 500.00},
		{Owner: "James", Balance: 100.00},
		{Owner: "John", Balance: 1900.00},
	}

	err := accounts[0].Deposit(100)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println(accounts[0].Balance)
	}

	err2 := accounts[1].Deposit(0)
	if err2 != nil {
		fmt.Println("Error: ", err2)
		//return пока не завершаем програму
	} else {
		fmt.Println(accounts[1].Balance)
	}

	err3 := accounts[2].Deposit(-10)
	if err3 != nil {
		fmt.Println("Error: ", err3)
		//return не завершаем
	} else {
		fmt.Println(accounts[2].Balance)
	}

	err4 := accounts[1].Withdraw(200)
	if err4 != nil {
		fmt.Println("Error: ", err4)
		//return
	} else {
		fmt.Println(accounts[1].Balance)
	}

	err5 := accounts[0].Withdraw(1000)
	if err5 != nil {
		fmt.Println("Error: ", err5)
		//return
	} else {
		fmt.Println(accounts[0].Balance)
	}

	err6 := accounts[2].Withdraw(50)
	if err6 != nil {
		fmt.Println("Error: ", err6)
		//return
	} else {
		fmt.Println(accounts[2].Balance)
	}

	err7 := accounts[1].Withdraw(0)
	if err7 != nil {
		fmt.Println("Error: ", err7)
		//return
	} else {
		fmt.Println(accounts[1].Balance)
	}

	fmt.Println(accounts[0].CanAfford(100))
	fmt.Println(accounts[2].CanAfford(2000))
	fmt.Println(accounts[1].CanAfford(0))
	fmt.Println(accounts[0].CanAfford(-1))

}

func (a *BankAccount) Deposit(amount float64) error {
	if amount > 0 {
		a.Balance += amount
	} else {
		return errors.New("deposit amount must be positive")
	}
	return nil
}

func (a *BankAccount) Withdraw(amount float64) error {
	if amount > 0 {
		if a.Balance >= amount {
			a.Balance -= amount
		} else {
			return errors.New("insufficient funds")
		}
	} else {
		return errors.New("withdraw amount must be positive")
	}
	return nil
}

func (a BankAccount) CanAfford(amount float64) bool {
	if amount > 0 && a.Balance >= amount {
		return true
	}

	return false
}

func (u *User) Birthday() {
	u.Age++
}

func (c *Counter) Increment() {
	c.Value++
}

func (c Counter) IncrementCopy() {
	c.Value++
}

func (c *Counter) Add(amount int) {
	c.Value += amount
}

func increment(value *int) {
	*value++
}
