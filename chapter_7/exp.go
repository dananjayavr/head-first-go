package main

import "fmt"

func main() {
	var votes []int

	for i := 0; i <= 4; i++ {
		votes = append(votes, 1)
	}

	fmt.Println(votes)
}
