package main

import (
	//    "fmt"
	"net/http"
	"text/template"
	//    "errors"
	"log"
)

func Home(writer http.ResponseWriter, reader *http.Request) {
	var templateHtml *template.Template
	templateHtml = template.Must(template.ParseFiles("main.html"))
	log.Println(templateHtml.Execute(writer, nil))
}

func main() {
	log.Println("Server started on: http://localhost:8000")
	//  var template_html *template.Template
	//template_html = template.Must(template.ParseFiles("main.html"))
	http.HandleFunc("/", Home)
	//  http.HandleFunc("/show", Show)
	//  http.HandleFunc("/new", New)
	//  http.HandleFunc("/edit", Edit)
	//  http.HandleFunc("/insert", Insert)
	//  http.HandleFunc("/update", Update)
	//    http.HandleFunc("/delete", Delete)
	log.Println(http.ListenAndServe(":8000", nil))
}
