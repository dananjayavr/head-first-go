package main

import (
	"dates"
	"fmt"
	"greeting/deutsch"
)

func main()  {
	days := 3
	deutsch.Hallo()
	fmt.Println("Your appointment is in", days, "days.")
	fmt.Println("With a follow-up in", days + dates.DaysInWeek, "days.")
	deutsch.GutenTag()
}
