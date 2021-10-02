package main

import (
	"fmt"
	"github.com/dananjayavr/calendar"
	"log"
)

func main() {
	date := calendar.Date{}

	err := date.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(24)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)

	wrongDate := calendar.Date{}
	err = wrongDate.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(wrongDate.Day())

	event := calendar.Event{}
	err = event.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(24)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetTitle("It' Christmas!")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event)
}
