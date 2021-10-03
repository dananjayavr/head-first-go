package main

import (
	"fmt"
	"math"
)

func myFunc(params ...int) int  {
	var sum int = 0
	for _, param := range params {
		sum += param
	}

	return sum
}

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64  {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}

	return  result
}

func average(numbers ...float64) float64  {
	numCount := float64(len(numbers))
	var sum float64
	for _, i2 := range numbers {
		sum += i2
	}

	return sum / numCount
}

func main() {
	fmt.Println(myFunc(3,4,2,4,5))
	fmt.Println(myFunc(3))
	fmt.Println(myFunc(3,4,2,4))
	fmt.Println(myFunc(3,4,2,4,3,4,1,3,5))

	fmt.Println(maximum(71.8,56.2,89.5))
	fmt.Println(maximum(90.7,89.7,98.5,92.3))

	fmt.Println(inRange(1,100,-12.5,3.2,0,50,103.5))
	fmt.Println(inRange(-10,10,4.1,12,-12,-5.2))

	fmt.Println(average(54.3,42.3,54.1,99.2,13.2,565.3))
	fmt.Println(average(54.3))
}
