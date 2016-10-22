package main

import (
	"bytes"
	"io"
	"os"
	"text/template"
)

// rnS OMIT
func render(w io.Writer, layout *template.Template,
	t *template.Template, tname string, tdata interface{}) {
	// eTS OMIT
	buf := &bytes.Buffer{}
	t.ExecuteTemplate(buf, tname, nil)
	// eTE OMIT
	layout.Funcs(template.FuncMap{ // HL
		"body": func() string {
			return buf.String()
		},
	})
	layout.Execute(w, nil)
}

// rnE OMIT

func main() {
	// layS OMIT
	layout := template.Must(template.New("layout").
		Funcs(template.FuncMap{"body": func() string {
			return "I am a layout."
		}}).
		Parse(`LetterHead
Dear
{{ body }}
Sincerely.
`))
	// layE OMIT
	// tS OMIT
	t := template.Must(template.New("").Parse(`
{{ define "gn" }}
I am pleased ...
{{ end }}
{{ define "bn" }}
We are sorry ...
{{ end }}
`))
	// tE OMIT

	// mnS OMIT
	render(os.Stdout, layout, t, "gn", nil)

	// mnE OMIT
}
