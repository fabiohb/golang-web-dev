package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", pics)
	http.Handle("/pics/", http.FileServer(http.Dir("./starting-files/public")))
	http.ListenAndServe(":8080", nil)
}

func pics(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("ERROS: ", err)
	}
}
