package main

import (
	"os"
	"text/template"
)

func main() {
	// S OMIT
	t := template.Must(template.New("ab").Parse(
		`{{- define "a"}}I am A for "{{ .aVal }}".
Still in A: {{template "b" . }}{{end}}
{{- define "b"}}I am B for "{{ .bVal }}".{{end}}
{{- template "a" . }} 
`))
	t.Execute(os.Stdout, map[string]string{"aVal": "apple", "bVal": "boy"})
	// E OMIT
}
