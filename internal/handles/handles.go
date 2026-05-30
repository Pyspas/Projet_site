package main

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Home")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")

	if err != nil {
		//fmt.Println("error")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
