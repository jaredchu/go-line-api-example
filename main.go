package main

import (
	"net/http"
	"html/template"
)

func init() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request)  {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)
}
