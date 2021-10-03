package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var mySlice []string
	mySlice = make([]string, 7)

	mySlice[0] = "do"
	mySlice[1] = "re"
	mySlice[2] = "mi"

	fmt.Println(mySlice[0])
	fmt.Println(mySlice[1])

	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])

	fmt.Println(len(primes))
	fmt.Println(len(mySlice))

	letters := make([]string, 5)

	letters[0] = "A"
	letters[1] = "B"
	letters[2] = "C"
	letters[3] = "D"
	letters[4] = "E"

	for i, letter := range letters {
		fmt.Println(i, "=>", letter)
	}

	notes := []string{"do","re","mi","fa","so","la","ti"}
	fmt.Println(notes[3],notes[6],notes[0])
	myPrimes := []int{
		2,
		3,
		5,
	}
	for i, prime := range myPrimes {
		fmt.Println(i, "=>", prime)
	}

	underlyingArray := [5]string{"a","b","c","d","e"}
	slice1 := underlyingArray[1:len(underlyingArray)]
	slice2 := underlyingArray[:len(underlyingArray)]
	fmt.Println(slice1)
	fmt.Println(slice2)

	// Appending an' that
	fmt.Println("===============================")
	slice := []string{"a","b"}
	fmt.Println(slice, len(slice))
	slice = append(slice,"c")
	fmt.Println(slice,len(slice))
	slice = append(slice,"d","e")
	fmt.Println(slice, len(slice))
}
