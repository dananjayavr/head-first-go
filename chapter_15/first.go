package main

import "fmt"

func sayHi()  {
	fmt.Println("Hi!")
}

func sayBye()  {
	fmt.Println("Bye!")
}

func twice(theFunction func())  {
	theFunction()
	theFunction()
}

func divide(a int, b int) float64  {
	return float64(a) / float64(b)
}

func callFunction(passedFunction func())  {
	passedFunction()
}

func callTwice(passedFunction func())  {
	passedFunction()
	passedFunction()
}

func callWithArguments(passedFunction func(s string, t bool))  {
	passedFunction("This sentence is", false)
}

func printReturnValue(passedFunction func() string)  {
	fmt.Println(passedFunction())
}

func functionA()  {
	fmt.Println("function called")
}
func functionB() string {
	fmt.Println("function called")
	return "returning from function"
}
func functionC(a string, b bool)  {
	fmt.Println("function called")
	fmt.Println(a,b)
}

func main() {
	var myFunction func()
	var myOtherFunction func(int, int) float64
	myFunction = sayHi
	myOtherFunction = divide
	fmt.Println(myOtherFunction(2,4))
	myFunction()
	twice(myFunction)
	twice(sayBye)

	callFunction(functionA)
	callTwice(functionA)
	callWithArguments(functionC)
	printReturnValue(functionB)
}
