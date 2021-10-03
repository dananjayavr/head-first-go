package main

import (
	"fmt"

	"github.com/dananjayavr/magazine"
)

func main() {
	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)

	subscriber := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}

	fmt.Println(subscriber.Name, subscriber.Rate)

	var employee magazine.Employee
	employee.Name = "Wayne Carr"
	employee.Salary = 60000

	fmt.Println(employee.Name, employee.Salary)

	var address magazine.Address
	address.Street = "123 Oak St"
	address.City = "Omaha"
	address.State = "NE"
	address.PostalCode = "68111"
	fmt.Println(address)

	employee.HomeAddress = address

	fmt.Println(employee)

	// Both forms below are equivalent

	// subscriber.Address.City = "Portland"
	// subscriber.Address.Street = "456 Elm St"
	// subscriber.Address.State = "OR"
	// subscriber.Address.PostalCode = "97222"

	subscriber.City = "Portland"
	subscriber.Street = "456 Elm St"
	subscriber.State = "OR"
	subscriber.PostalCode = "97222"

	fmt.Println(subscriber)
}
