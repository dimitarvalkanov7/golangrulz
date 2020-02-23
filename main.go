package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", ":8082", "http service address")
	tmpl = template.Must(template.New("tmpl").Parse(`
		<!DOCTYPE html><html><body><center>
		<h1>GoLang Rulz!</h1>
		</center></body></html>`))
)

func index(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(*addr, nil)
}
