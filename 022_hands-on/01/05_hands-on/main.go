package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.Handle("/", http.HandlerFunc(i))
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/me/", http.HandlerFunc(m))

	http.ListenAndServe(":8080", nil)
}

func i(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>INDEX</h1>")
}

func d(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>DOG</h1>")
}

func m(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(w, "me.gohtml", "Fabio")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
