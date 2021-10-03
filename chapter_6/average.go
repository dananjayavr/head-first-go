// average calculates the average of several numbers
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64  {
	var sum float64
	for _, f := range numbers {
		sum += f
	}

	return sum / float64(len(numbers))
}

func main() {
	arguments := os.Args[1:]
	var numbers []float64
	if len(arguments) > 0 {
		for _, argument := range arguments {
			parsed, err := strconv.ParseFloat(argument, 64)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, parsed)
		}
		/*for _, argument := range arguments {
			number, err := strconv.ParseFloat(argument, 64)
			if err != nil {
				log.Fatal(err)
			}
			sum += number
		}
		sampleCount := float64(len(arguments))*/
		fmt.Printf("Average: %0.2f\n", average(numbers...))
	} else {
		fmt.Println("Hint: Supply at least one float64 argument.")
	}
}
