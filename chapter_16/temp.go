package main

import (
	"html/template"
	"log"
	"os"
)

func myCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{})  {
	tmpl, err := template.New("test").Parse(text)
	myCheck(err)
	err = tmpl.Execute(os.Stdout, data)
	myCheck(err)
}

func main() {
/*	text := "Here's my template!\nAction: {{.}}\nTemplate end\n"
	tmpl, err := template.New("test").Parse(text)
	myCheck(err)
	err = tmpl.Execute(os.Stdout, "ABC")
	err = tmpl.Execute(os.Stdout, "42")
	err = tmpl.Execute(os.Stdout, true)
	myCheck(err)

	executeTemplate("start {{if .}}Dot is true!{{end}}\nDot is: {{.}}\n",true)
	executeTemplate("start {{if .}}Dot is true!{{end}}\nDot is: {{.}}\n",false)
	executeTemplate("Dot is: {{.}}\n",123.6)*/

	templateText := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n"
	executeTemplate(templateText, []string{"do","re","mi"})

	pricesTemplate := "Prices:\n{{range .}}${{.}}\n{{end}}"
	executeTemplate(pricesTemplate,[]float64{1.25,0.99,27})

	type Part struct {
		Name string
		Count int
	}

	type Subscriber struct {
		Name string
		Rate float64
		Active bool
	}

	structTemplate := "Name: {{.Name}}\nCount:{{.Count}}\n"
	executeTemplate(structTemplate,Part{
		Name: "Fuses",
		Count: 5,
	})
	executeTemplate(structTemplate,Part{
		Name: "Cables",
		Count: 2,
	})

	subscriberTemplate := "Name: {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	subscriber := Subscriber{
		Name: "John Doe",
		Rate: 4.99,
		Active: true,
	}
	executeTemplate(subscriberTemplate,subscriber)
	anotherSubscriber := Subscriber{
		Name: "Wayne Carr",
		Rate: 5.99,
		Active: false,
	}
	executeTemplate(subscriberTemplate,anotherSubscriber)
}
