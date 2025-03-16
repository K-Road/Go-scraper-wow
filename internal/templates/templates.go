package templates

import "html/template"

func LoadTemplates() (*template.Template, error) {
	return template.ParseFiles("internal/templates/base.html", "internal/templates/content-score.html")
}
