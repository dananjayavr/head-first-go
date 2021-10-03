package main

import "fmt"

func main() {
	var quantity int
	var length, width float64
	var customerName string
	//var dimensions float64

	quantity = 42
	length = 0.0
	width = 0.0

	customerName = "Alan Partridge"
	quantity = int(length * width)

	fmt.Println(customerName)
	fmt.Println(quantity)

}
