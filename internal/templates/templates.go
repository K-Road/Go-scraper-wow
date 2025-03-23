package templates

import (
	"html/template"
	"time"
)

func LoadTemplates() (*template.Template, error) {
	return template.ParseFiles("internal/templates/base.html", "internal/templates/content-score.html")
}

// Data struct for the HTML template
type BaseTemplateData struct {
	Date time.Time
}

func GetBaseTemplateData() BaseTemplateData {
	baseTemplateData := BaseTemplateData{
		Date: time.Now(),
	}
	return baseTemplateData
}

type PageData struct {
	BaseData  BaseTemplateData
	ScoreData ScoreContentData
}

type ScoreContentData struct {
	Name       string
	Score      float64
	ClassColor string
}
