package main

import (
	"os"
	"text/template"
	"time"
)

func main() {
	// F_S OMIT
	t := template.Must(template.New("hello").Parse(`Hello {{ printf "%s%s" .salut .name }}.
The time is {{ call .time }}.
In {{ .offset }} hour(s) it will be {{ call .toffset .offset }}.
`))
	t.Execute(os.Stdout, map[string]interface{}{"salut": "Mr",
		"name": "Tan", "time": func() string {
			return time.Now().Format("3:04:05PM")
		},
		"offset": 2,
		"toffset": func(i int) string {
			return time.Now().Add(time.Duration(i) * time.Hour).
				Format("3:04:05PM")
		}})
	// F_E OMIT
}
