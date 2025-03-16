package main

import (
	"fmt"
	"os"
	"time"

	"github.com/K-Road/Go-scraper-wow/internal/scraper"
	"github.com/K-Road/Go-scraper-wow/internal/templates"
)

// Data struct for the HTML template
type TemplateData struct {
	Name       string
	Score      float64
	Date       time.Time
	ClassColor string
}

// const baseHTMLTemplate = `
// <!DOCTYPE html>
// <html>
// <head>
// 	<title>WOW Character Report</title>
// 	<style>
// 		body {font-family: Ariel, sans-serif; padding: 20px; }
// 		h2 { color: {{.ClassColor}}; }
// 	</style>
// </head>
// <body>
// 	<h1>Wow Character Data</h1>
// 	{{ template "contentScore" . }}
// 		<p>{{.Date}}</p>
// </body>
// </html>`

// const contentScoreHTMLTemplate = `
// {{define "contentScore"}}
// <h2>{{.Name}}</h2>
// <p>MPlus Score: {{.Score}}</p>
// {{end}}`

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
		Name:       data.Name,
		Score:      score,
		Date:       time.Now(),
		ClassColor: getClassColor(data.Class),
	}

	tmpl, err := templates.LoadTemplates()
	if err != nil {
		return "", err
	}
	//	tmpl, err := template.New("base").Parse(baseHTMLTemplate)
	// tmpl, err := template.ParseFiles("template/base.html", "content-score.html")
	// if err != nil {
	// 	return "", err
	// }
	// _, err = tmpl.New("contentScore").Parse(contentScoreHTMLTemplate)
	// if err != nil {
	// 	return "", err
	// }

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
	classColors := map[string]string{
		"Warrior":      "#C79C6E",
		"Paladin":      "#F58CBA",
		"Hunter":       "#ABD473",
		"Rogue":        "#FFF569",
		"Priest":       "#FFFFFF",
		"Death Knight": "#C41F3B",
		"Shaman":       "#0070DE",
		"Mage":         "#69CCF0",
		"Warlock":      "#9482C9",
		"Monk":         "#00FF96",
		"Druid":        "#FF7D0A",
		"Demon Hunter": "#A330C9",
	}
	color, exists := classColors[class]
	if !exists {
		return "#FFFFFF" // Default to white if class not found
	}
	return color
}
