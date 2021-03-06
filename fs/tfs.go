package main

import (
	"os"
	"text/template"
)

func main() {
	// FS_S OMIT
	t := template.Must(template.New("tplAB").ParseGlob("fs/*.tpl"))
	m := map[string]interface{}{"aVal": "apple",
		"bVal": "boy"}
	t.Execute(os.Stdout, m)
	t.ExecuteTemplate(os.Stdout, "tplB", m["bVal"])
	// FS_E OMIT
}
