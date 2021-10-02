package main

import "fmt"

type MyNumber int

func (n *MyNumber) Display() {
	fmt.Println(*n)
}
func (n *MyNumber) Double() {
	*n *= 2
}

func main() {
	number := MyNumber(4)
	number.Double()
	number.Display()
}
