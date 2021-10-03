package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func sayHi() {
	fmt.Println("Hi!")
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

func main() {
	//sayHi()
	//repeatLine("Don't Panic!", 42)
	//var amount float64
	//amount = calculatePaintNeeded(4.2, 3.0)
	//fmt.Printf("%#v\n", formattedString)
	//fmt.Printf("%.2f liters needed\n", amount)
	//amount += calculatePaintNeeded(5.2, 3.5)
	//fmt.Printf("%.2f liters needed\n", amount)
	//amount += calculatePaintNeeded(5.0, 3.3)
	//fmt.Printf("%.2f\n", amount)

	/*	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
		fmt.Println("------------------------------------")

		fmt.Printf("%12s | %2d\n", "Stamps", 50)
		fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
		fmt.Printf("%12s | %2d\n", "Tape", 99)*/

	//fmt.Printf("Total: %0.2f liters\n", amount)


	amount, err := calculatePaintNeeded(4.2,3.0)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f liters needed\n", amount)
	}

	quotient, err := divide(5.6,2.0)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Quotient: %0.2f\n", quotient)
	}

	d := 2
	fmt.Println("Before:",d)
	double(&d)
	fmt.Println("After:", d)

	fmt.Println(d)
	fmt.Println(&d)
	fmt.Println(reflect.TypeOf(&d))

	var myInt int
	var myIntPointer *int

	myInt = 42
	myIntPointer = &myInt

	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer << 2)
	fmt.Println(myInt)
	*myIntPointer = 43
	fmt.Println(myInt)

	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)

}

func calculatePaintNeeded(width float64, height float64) (float64, error) {
	var area float64
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid.", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid.", height)
	}
	area = width * height
	return area / 10.0,nil
}

func divide(dividend float64, divisor float64) (float64, error)  {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return dividend / divisor, nil
}

func double(number *int) *int {
	*number *= 2
	return number
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}