package main

import (
	"fmt"
	"mypkg"
)

func main() {
	var value mypkg.MyInterface
	value = mypkg.MyType(5) // possible because MyType *satisfied* MyInterface
	value.MethodWithoutParameters()
	value.MethodWithParameters(127.3)
	fmt.Println(value.MethodWithReturnValue())
}
