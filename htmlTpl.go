package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
)

// LAY_S OMIT
var layoutFuncs = template.FuncMap{
	"yield": func() (string, error) {
		return "", errors.New("bad yield call")
	},
}
var layout = template.Must(template.
	New("layout.html").
	Funcs(layoutFuncs).
	ParseFiles("html/tpl/layout.html"),
)
var myTemplates = template.Must(template.New("").ParseGlob("html/tpl/**/*.html"))

// LAY_E OMIT

// RDR_S OMIT

// RenderTemplate This demo of RenderTemplate just renders to os.Stdout.
// It would typically write to w of w http.ResponseWriter
// and perhaps augment data from r of r http.Request
// thus: RenderTemplate(w http.ResponseWriter, r http.Request, name string, data interface{})
func RenderTemplate(name string, data interface{}) {
	funcs := template.FuncMap{
		// 100 OMIT
		"yield": func() (template.HTML, error) {
			buf := &bytes.Buffer{}
			err := myTemplates.ExecuteTemplate(buf, name, data)
			return template.HTML(buf.String()), err // treats buf string as trusted HTML // HL
		},
		// 200 OMIT
	}
	layoutClone, _ := layout.Clone() // Work on clone of layout to ensure
	layoutClone.Funcs(funcs)         // goroutines do not step on each other.
	err := layoutClone.Execute(os.Stdout, data)
	if err != nil {
		fmt.Printf("something went wrong:\n%s", err)
	}
}

// RDR_E OMIT

func main() {
	// CALL_S OMIT
	RenderTemplate("index/home", map[string]interface{}{
		"myNav":      "my navigation",
		"myMainBody": "my main body",
		"myFooter":   "my footer",
	})
	// CALL_E OMIT
}
