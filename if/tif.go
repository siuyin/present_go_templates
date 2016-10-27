package main

import (
	"os"
	"text/template"
)

func main() {
	// IF_S OMIT
	t := template.Must(template.New("hello").Parse(`{{if .Turn -}}
Turning: {{ .Turn }}. {{- else -}}
Going straight.
{{- end }}
`))
	t.Execute(os.Stdout, struct{ Turn string }{""})
	t.Execute(os.Stdout, struct{ Turn string }{"Right"})
	t.Execute(os.Stdout, struct{ Turn bool }{true})
	t.Execute(os.Stdout, struct{ Turn int }{0})
	// IF_E OMIT
}
