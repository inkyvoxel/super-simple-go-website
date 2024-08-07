package app

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := filepath.Join("web", "template", "index.html")
	template, err := template.ParseFiles(templatePath)

	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	err = template.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}
