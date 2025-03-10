package main

import (
	"fmt"

	"github.com/K-Road/Go-scraper-wow/config"
	"github.com/K-Road/Go-scraper-wow/internal/gcs"
	"github.com/K-Road/Go-scraper-wow/internal/scraper"
)

type CharacterConfig struct {
	CharacterName string `json:"character_name"`
	Realm         string `json:"realm"`
	Region        string `json:"region"`
}

// type MythicPlusScore struct {
// 	Scores struct {
// 		DPS float64 `json:"dps"`
// 	} `json:"scores"`
// }

// type APIResponse struct {
// 	MythicPlusScoresBySeason []MythicPlusScore `json:"mythic_plus_scores_by_season"`
// }

func main() {
	//Load GSC
	gcsConfig, err := config.LoadGCSConfig()
	if err != nil {
		fmt.Println("Error loading GCS", err)
		return
	}

	//load characterconfig
	cconfig, err := config.LoadCharacterConfig("character.config.json") //TODO change to dynamic list
	if err != nil {
		fmt.Println("Error loading character config:", err)
		return
	}
	fmt.Printf("Scraping %s\n", cconfig.CharacterName)

	charData, err := scraper.ScrapeCharacterData(cconfig.CharacterName, cconfig.Realm, cconfig.Region)
	if err != nil {
		fmt.Printf("Error scraping %s data:", cconfig.CharacterName)
		return
	}

	if len(charData.MythicPlusScoresBySeason) > 0 {
		fmt.Println("Score:", charData.MythicPlusScoresBySeason[0].Scores.Dps)
	} else {
		fmt.Println("No Mythic+ score found")
	}

	filename, err := GenerateHTMLReport(cconfig.CharacterName)
	if err != nil {
		fmt.Println("Error generating report:", err)
		return
	}

	err = gcs.UploadtoGCS(gcsConfig.BucketName, filename)
	if err != nil {
		fmt.Println("Error uploading", err)
		return
	}

}
