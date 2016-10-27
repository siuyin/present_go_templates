package main

import (
	"os"
	"text/template"
	"time"
)

// S_S OMIT
func sendInts() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch) // HL
	}()
	return ch
}

// S_E OMIT
func main() {
	// R_S OMIT
	t := template.Must(template.New("hello").Parse(`{{range . -}}
{{ . }} {{ end }}
`))
	queue := sendInts()
	t.Execute(os.Stdout, queue) // t does a <-queue
	// R_E OMIT
}
