package main

import (
	"html/template"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Invoice struct {
	Name string
	Paid bool
	Charges []float64
	Total float64
}

func main() {
	html, err := template.ParseFiles("bill.html")
	checkError(err)
	bill := Invoice{
		Name: "Mary Gibbs",
		Paid: true,
		Charges: []float64{23.19,1.13,42.79},
		Total: 67.11,
	}
	err = html.Execute(os.Stdout, bill)
	checkError(err)
}
