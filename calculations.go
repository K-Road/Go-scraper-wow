package main

import "github.com/K-Road/Go-scraper-wow/internal/scraper"

func mplusScore(data scraper.APIResponse) float64 {
	var score float64
	if len(data.MythicPlusScoresBySeason) > 0 {
		score = data.MythicPlusScoresBySeason[0].Scores.Dps
	} else {
		score = 0.0
	}
	return score
}
