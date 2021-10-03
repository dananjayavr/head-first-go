package main

import "fmt"

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	fmt.Printf("%#v\n", myStruct)

	myStruct.number = 3.14
	fmt.Println(myStruct.number)

	myStruct.word = "pi"
	myStruct.toggle = true

	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)

	type subscriber struct {
		name   string
		rate   float64
		active bool
	}

	var sub subscriber

	sub.name = "John Doe"
	sub.rate = 4.99
	sub.active = true

	fmt.Println("Name: ", sub.name)
	fmt.Println("Monthly rate: ", sub.rate)
	fmt.Println("Active? ", sub.active)
}
