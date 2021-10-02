package main

import "fmt"

func greeting(myChannel chan string)  {
	myChannel <- "hi"
}

func abc(channel chan string)  {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func cde(channel chan string)  {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
/*	myChannel := make(chan string)
	go greeting(myChannel)
	fmt.Println(<- myChannel)*/
	channel1 := make(chan string)
	channel2 := make(chan string)
	go abc(channel1)
	go cde(channel2)
	fmt.Println(<- channel1)
	fmt.Println(<- channel2)
	fmt.Println(<- channel1)
	fmt.Println(<- channel2)
	fmt.Println(<- channel1)
	fmt.Println(<- channel2)
	fmt.Println()

}
