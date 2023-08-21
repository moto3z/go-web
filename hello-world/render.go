package main

import (
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	pt, _ := template.ParseFiles("./templates/" + tmpl)
	err := pt.Execute(w, nil)
	if err != nil {
		print("error")
	}
}
