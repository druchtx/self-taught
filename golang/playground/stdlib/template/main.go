package main

import (
	"html/template"
	"os"
)

type Inventory struct {
	Material string
	Count    int
}

func main() {
	sweaters := []Inventory{
		{"wool", 100},
		{"polyester", 50},
	}

	tpl, err := template.New("test").Parse("{{range .}} {{.Material }} {{.Count }}  \n{{end}}")
	if err != nil {
		return
	}

	if err := tpl.Execute(os.Stdout, sweaters); err != nil {
		return
	}

}
