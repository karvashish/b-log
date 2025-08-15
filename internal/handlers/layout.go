package handlers

import (
	"html/template"
	"net/http"
)

var layoutTmpl *template.Template

func SetLayoutTemplate(t *template.Template) {
	layoutTmpl = t
}

type layoutData struct {
	Title   string
	Content template.HTML
}

func renderLayout(w http.ResponseWriter, title, content string) {
	if layoutTmpl == nil {
		http.Error(w, "layout template not initialized", http.StatusInternalServerError)
		return
	}
	data := layoutData{
		Title:   title,
		Content: template.HTML(content),
	}

	if err := layoutTmpl.ExecuteTemplate(w, "layout.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
