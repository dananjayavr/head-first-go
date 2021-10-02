package main

import "fmt"

type MyType string
type Number int

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

func (n *Number) Double() {
	*n *= 2
}

func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func (m MyType) WithReturn() int {
	return len(m)
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func main() {
	value := MyType("a MyType Value")
	value.sayHi()
	anotherValue := MyType("another value")
	anotherValue.sayHi()
	value.MethodWithParameters(4, true)

	fmt.Println(value.WithReturn())
	fmt.Println(anotherValue.WithReturn())

	number := Number(4)
	fmt.Println(number)
	number.Double()
	fmt.Println(number)

	value.pointerMethod()
	value.method()
}
