package main

import (
	"os"
	"text/template"
)

func main() {
	// tS OMIT
	t := template.Must(template.New("ab").Parse(`
{{- define "a"}}I am A for "{{ .aVal }}".
  Still in A: {{ template "b" .bStuff }}{{end}}
{{- define "b"}}I am B for "{{ .bVal }}".{{end}}

{{- template "a" . }}
`))
	t.Execute(os.Stdout, map[string]interface{}{
		"aVal":   "apple",
		"bStuff": map[string]interface{}{"bVal": "boy"},
	})
	// tE OMIT
}
