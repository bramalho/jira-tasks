package controller

import (
	"html/template"
	"net/http"

	"github.com/bramalho/jira-tasks/model"
	"github.com/bramalho/jira-tasks/service"
)

// HomeHandler controller method
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/template/index.html"))
	tmpl.Execute(w, struct {
		Users []*model.User
	}{
		service.New(),
	})
}
