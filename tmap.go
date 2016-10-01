package main

import (
	"os"
	"text/template"
)

func main() {
	// MP_S OMIT
	t := template.Must(template.New("hello").Parse(`Hello {{ .salut -}} 
	{{ .name }}. You have won {{ .qty }} apples!
`))
	t.Execute(os.Stdout, map[string]interface{}{"salut": "Mr",
		"name": "Tan", "qty": 3})
	// MP_E OMIT
}
