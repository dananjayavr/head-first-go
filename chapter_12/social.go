package main

import (
	"fmt"
	"log"
)

func Socialize() error {
	defer fmt.Println("Goodbye!") // will execute when all the other functions in Socialize() are complete.
	fmt.Println("Hello!")
	return fmt.Errorf("i dont't want to talk")
	//return nil
	//fmt.Println("Nice weather, eh?")
}

func main() {
	err := Socialize()
	if err != nil {
		log.Fatal(err)
	}
}
