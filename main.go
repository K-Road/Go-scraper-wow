package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type CharacterConfig struct {
	CharacterName string `json:"character_name"`
	Realm         string `json:"realm"`
	Region        string `json:"region"`
}

type MythicPlusScore struct {
	Scores struct {
		DPS float64 `json:"dps"`
	} `json:"scores"`
}

type APIResponse struct {
	MythicPlusScoresBySeason []MythicPlusScore `json:"mythic_plus_scores_by_season"`
}

func loadCharacterConfig(filename string) (CharacterConfig, error) {
	var cconfig CharacterConfig
	file, err := os.Open(filename)
	if err != nil {
		return cconfig, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cconfig)
	return cconfig, err
}

func scrapeCharacterData(characterName, realm, region string) {
	return
}

func main() {
	//load characterconfig
	cconfig, err := loadCharacterConfig("character.config.json") //TODO change to dynamic list
	if err != nil {
		fmt.Println("Error loading character config:", err)
		return
	}
	scrapeCharacterData(cconfig.CharacterName, cconfig.Realm, cconfig.Region)
	fmt.Printf("Scraping %s\n", cconfig.CharacterName)
}
