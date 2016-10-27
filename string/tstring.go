package main

import (
	"os"
	"text/template"
)

func main() {
	// STR_S OMIT
	t := template.Must(template.New("hello").Parse(`Hello {{ . }}!
--The End--`))
	t.Execute(os.Stdout, "world")
	// STR_E OMIT
}
