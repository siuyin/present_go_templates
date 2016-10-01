package main

import (
	"os"
	"text/template"
)

func main() {
	// S OMIT
	t := template.Must(template.New("ab").Parse(
		`{{define "a"}}I am A for "{{ . }}".{{end}}
{{- define "b"}}I am B for "{{ . }}".{{end}}
{{- template "a" .aVal }} {{template "b" .bVal }}
`))
	t.Execute(os.Stdout, map[string]string{"aVal": "apple", "bVal": "boy"})
	// E OMIT
}
