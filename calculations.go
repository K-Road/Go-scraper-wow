package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/K-Road/Go-scraper-wow/internal/scraper"
)

func mplusScore(data scraper.APIResponse) float64 {
	var score float64
	if len(data.MythicPlusScoresBySeason) > 0 {
		score = data.MythicPlusScoresBySeason[0].Scores.Dps
	} else {
		score = 0.0
	}
	return score
}

func gear(data scraper.APIResponse) {
	gear := data.Gear

	//Extract items
	items, err := extractItemsToMap(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Item level: ", gear.ItemLevelEquipped)

	//iterate over items to print ##TODO store for item comparisons
	for slot, item := range items {
		if itemMap, ok := item.(map[string]interface{}); ok {
			if itemID, ok := itemMap["item_id"]; ok {
				fmt.Printf("Slot: %s, Item: %v\n", slot, itemID)
			}
		}
	}
}

func extractItemsToMap(data scraper.APIResponse) (map[string]interface{}, error) {
	itemsJSON, err := json.Marshal(data.Gear.Items)
	if err != nil {
		return nil, err
	}

	var itemsMap map[string]interface{}
	err = json.Unmarshal(itemsJSON, &itemsMap)
	if err != nil {
		return nil, err
	}

	return itemsMap, nil
}
