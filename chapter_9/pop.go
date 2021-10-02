package main

import "fmt"

type Population int

func main() {
	var population Population
	population = Population(572)
	fmt.Println("Sleepy Creek county population:", population)
	fmt.Println("Congratulations, Kevin and Anna! It's a girl!")
	population += 1
	fmt.Println("Sleepy Creek county population:", population)
}
