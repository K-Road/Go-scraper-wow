package main

import (
	"fmt"
	"os"
	"time"

	"github.com/K-Road/Go-scraper-wow/internal/scraper"
	"github.com/K-Road/Go-scraper-wow/internal/templates"
)

// Data struct for the HTML template
// type TemplateData struct {
// 	Name       string
// 	Score      float64
// 	ClassColor string
// }

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

	score := mplusScore(data)

	baseData := templates.GetBaseTemplateData()
	//Generate template data //##TODO pull this out. Put in template data generator. Pass into this func as param
	scoreData := templates.ScoreContentData{
		Name:       data.Name,
		Score:      score,
		ClassColor: getClassColor(data.Class),
	}

	pageData := templates.PageData{
		BaseData:  baseData,
		ScoreData: scoreData,
	}

	tmpl, err := templates.LoadTemplates()
	if err != nil {
		return "", err
	}

	fileName := "/tmp/output.html"
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = tmpl.Execute(file, pageData)
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
