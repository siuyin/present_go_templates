package main

import (
	"os"
	"text/template"
)

func main() {
	// STR_S OMIT
	t := template.Must(template.
		New("hello").
		Parse(`Hello {{ .Salutation -}} {{ .Name }}!
Is your age {{ .Age }}?`))
	t.Execute(os.Stdout, struct {
		Salutation, Name string
		Age              int
	}{"Mr", "Tan", 42})
	// STR_E OMIT
}
