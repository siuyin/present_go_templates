package main

import (
	"os"
	"text/template"
)

func main() {
	// R_S OMIT
	t := template.Must(template.New("hello").Parse(`{{len .Dat}} items.
{{range $i,$val := .Dat -}}
	Displaying {{ $i }}: {{ $val }} | {{ . }}.
{{ else -}}
-- no data --
{{- end }}
`))
	t.Execute(os.Stdout, struct{ Dat []string }{[]string{"a", "b", "", "d"}})

	t.Execute(os.Stdout, struct{ Dat []string }{})
	// R_E OMIT
}
