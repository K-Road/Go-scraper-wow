package main

import (
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/K-Road/Go-scraper-wow/internal/scraper"
)

// Data struct for the HTML template
type TemplateData struct {
	Name  string
	Score float64
	Date  time.Time
}

const htmlTemplate = `
		<!DOCTYPE html>
		<html>
		<head>
			<title>WOW Character Report</title>
			<style>
				body {font-family: Ariel, sans-serif; padding: 20px; }
				h1 { color: ; #007bff}
				h2 { color: %s; }
			</style>
		</head>
		<body>
			<h1>Wow Character Data</h1>
			<h2>{{.Name}}</h2>
			<p>{{.Score}}</p>
			<br>
			<p>{{.Date}}</p>
		</html>`

func GenerateHTMLReport(data scraper.APIResponse) (string, error) {
	color := getClassColor(data.Class)

	htmlContent := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>WOW Character Report</title>
			<style>
				body {font-family: Ariel, sans-serif; padding: 20px; }
				h1 { color: ; #007bff}
				h2 { color: %s; }
			</style>
		</head>
		<body>
			<h1>Wow Character Data</h1>
			<h2>%s</h2>
			<p>%v</p>
			<br>
			<p>%v</p>
		</html>`, color, data.Name, data, time.Now())

	file, err := os.Create("report.html")
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = file.WriteString(htmlContent)
	if err != nil {
		return "", err
	}
	return file.Name(), err
}

func GenerateHTML(data scraper.APIResponse) (string, error) {

	var score float64
	if len(data.MythicPlusScoresBySeason) > 0 {
		score = data.MythicPlusScoresBySeason[0].Scores.Dps
	} else {
		score = 0.0
	}

	templateData := TemplateData{
		Name:  data.Name,
		Score: score,
		Date:  time.Now(),
	}

	tmpl, err := template.New("report").Parse(htmlTemplate)
	if err != nil {
		return "", err
	}
	fileName := "/tmp/output.html"
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = tmpl.Execute(file, templateData)
	return fileName, err
}

func getClassColor(class string) string {
	color := "#007bff"
	switch class {
	case "Paladin":
		color = "#F48CBA"
	case "Warlock":
		color = "#8788EE"
	}
	return color
}
