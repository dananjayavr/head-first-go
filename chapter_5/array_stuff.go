package main

import (
	"fmt"
	"time"
)

func main() {
	var notes [7]string

	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"

	fmt.Println(notes[0])
	fmt.Println(notes[1])

	var primes [5]int

	primes[0] = 2
	primes[1] = 3

	fmt.Println(primes[0])
	fmt.Println(primes[1])

	var dates [3]time.Time

	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	dates[0] = time.Unix(1508632200, 0)

	fmt.Println(dates[0])
	fmt.Println(dates[1])
	fmt.Println(dates[2])

	var counters [3]int

	counters[0]++

	for i := 0; i < 3; i++ {
		counters[i]++
		fmt.Println(counters[i])
	}

	var literalNotes = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}

	for i := 0; i < 7; i++ {
		fmt.Println(literalNotes[i])
	}

	literalPrimes := [5]int{2, 3, 5, 7, 11}

	for i, prime := range literalPrimes {
		fmt.Println(i, "==>", prime)
	}

	hp := [7]string{
		"Harry Potter and the Sorcerer's Stone",
		"Harry Potter and the Chamber of Secrets",
		"Harry Potter and  the Prisoner of Azkaban",
		"Harry Potter and the Goblet of Fire",
		"Harry Potter and the Order of Phoenix",
		"Harry Potter and the Half-Blood Prince",
		"Harry Potter and the Deathly Hallows",
	}

	for i := 0; i < len(hp); i++ {
		fmt.Println("Book", i, "=>", hp[i])
	}

	for _, book := range hp {
		fmt.Println(book)
	}

	//fmt.Printf("%#v",hp)
}
