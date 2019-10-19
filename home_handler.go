package main

import (
	"html/template"
	"net/http"
)

// HomeHandler controller method
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/template/index.html"))
	tmpl.Execute(w, struct {
		Users []*User
	}{
		GetUserData(),
	})
}
