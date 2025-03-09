package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MythicPlusScore struct {
	Scores struct {
		DPS float64 `json:"dps"`
	} `json:"scores"`
}

type APIResponse struct {
	MythicPlusScoresBySeason []MythicPlusScore `json:"mythic_plus_scores_by_season"`
}

func ScrapeCharacterData(characterName, realm, region string) {
	url := fmt.Sprintf("https://raider.io/api/v1/characters/profile?region=%s&realm=%s&name=%s&fields=gear%%2Cmythic_plus_scores_by_season:current",
		region, realm, characterName,
	)

	fmt.Println("Fetching data from:", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Recieved non-200 status code:", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var data APIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if len(data.MythicPlusScoresBySeason) > 0 {
		fmt.Println("Score:", data.MythicPlusScoresBySeason[0].Scores.DPS)
	} else {
		fmt.Println("No Mythic+ score found")
	}

}
