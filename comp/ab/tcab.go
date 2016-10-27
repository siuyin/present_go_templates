package main

import (
	"os"
	"text/template"
)

func main() {
	// tS OMIT
	t := template.Must(template.New("ab").Parse(
		`{{define "a"}}I am A for "{{ .aVal }}".{{end}}
		{{- define "b"}}I am B for "{{ .bVal }}" and "{{ .b2 }}".{{end}}

		{{- template "a" .aStuff }} {{template "b" .bStuff }}
`))
	t.Execute(os.Stdout, map[string]interface{}{
		"aStuff": map[string]string{"aVal": "apple"},
		"bStuff": map[string]string{"bVal": "boy", "b2": "bat"},
	})
	// tE OMIT
}
