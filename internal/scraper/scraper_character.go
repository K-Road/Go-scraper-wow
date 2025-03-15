package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type MythicPlusScore struct {
	Scores struct {
		DPS float64 `json:"dps"`
	} `json:"scores"`
}

type APIResponse struct {
	//MythicPlusScoresBySeason []MythicPlusScore `json:"mythic_plus_scores_by_season"`
	Name                     string    `json:"name"`
	Race                     string    `json:"race"`
	Class                    string    `json:"class"`
	ActiveSpecName           string    `json:"active_spec_name"`
	ActiveSpecRole           string    `json:"active_spec_role"`
	Gender                   string    `json:"gender"`
	Faction                  string    `json:"faction"`
	AchievementPoints        int       `json:"achievement_points"`
	ThumbnailURL             string    `json:"thumbnail_url"`
	Region                   string    `json:"region"`
	Realm                    string    `json:"realm"`
	LastCrawledAt            time.Time `json:"last_crawled_at"`
	ProfileURL               string    `json:"profile_url"`
	ProfileBanner            string    `json:"profile_banner"`
	MythicPlusScoresBySeason []struct {
		Season string `json:"season"`
		Scores struct {
			All    float64 `json:"all"`
			Dps    float64 `json:"dps"`
			Healer int     `json:"healer"`
			Tank   int     `json:"tank"`
			Spec0  int     `json:"spec_0"`
			Spec1  int     `json:"spec_1"`
			Spec2  float64 `json:"spec_2"`
			Spec3  int     `json:"spec_3"`
		} `json:"scores"`
		Segments struct {
			All struct {
				Score float64 `json:"score"`
				Color string  `json:"color"`
			} `json:"all"`
			Dps struct {
				Score float64 `json:"score"`
				Color string  `json:"color"`
			} `json:"dps"`
			Healer struct {
				Score int    `json:"score"`
				Color string `json:"color"`
			} `json:"healer"`
			Tank struct {
				Score int    `json:"score"`
				Color string `json:"color"`
			} `json:"tank"`
			Spec0 struct {
				Score int    `json:"score"`
				Color string `json:"color"`
			} `json:"spec_0"`
			Spec1 struct {
				Score int    `json:"score"`
				Color string `json:"color"`
			} `json:"spec_1"`
			Spec2 struct {
				Score float64 `json:"score"`
				Color string  `json:"color"`
			} `json:"spec_2"`
			Spec3 struct {
				Score int    `json:"score"`
				Color string `json:"color"`
			} `json:"spec_3"`
		} `json:"segments"`
	} `json:"mythic_plus_scores_by_season"`
	Gear struct {
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		Source            string    `json:"source"`
		ItemLevelEquipped float64   `json:"item_level_equipped"`
		ItemLevelTotal    int       `json:"item_level_total"`
		ArtifactTraits    int       `json:"artifact_traits"`
		Corruption        struct {
			Added     int           `json:"added"`
			Resisted  int           `json:"resisted"`
			Total     int           `json:"total"`
			CloakRank int           `json:"cloakRank"`
			Spells    []interface{} `json:"spells"`
		} `json:"corruption"`
		Items struct {
			Head struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Tier             string        `json:"tier"`
				Gems             []int         `json:"gems"`
				Enchants         []interface{} `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"head"`
			Neck struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Gems             []int         `json:"gems"`
				Enchants         []interface{} `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"neck"`
			Shoulder struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Tier             string        `json:"tier"`
				Gems             []interface{} `json:"gems"`
				Enchants         []interface{} `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"shoulder"`
			Back struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Enchant        int           `json:"enchant"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Gems             []interface{} `json:"gems"`
				Enchants         []int         `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"back"`
			Chest struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Enchant        int           `json:"enchant"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Tier             string        `json:"tier"`
				Gems             []interface{} `json:"gems"`
				Enchants         []int         `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"chest"`
			Waist struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Gems             []int         `json:"gems"`
				Enchants         []interface{} `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"waist"`
			Shirt struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Gems             []interface{} `json:"gems"`
				Enchants         []interface{} `json:"enchants"`
				Bonuses          []interface{} `json:"bonuses"`
			} `json:"shirt"`
			Wrist struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Enchant        int           `json:"enchant"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Gems             []int         `json:"gems"`
				Enchants         []int         `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"wrist"`
			Hands struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Gems             []interface{} `json:"gems"`
				Enchants         []interface{} `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"hands"`
			Legs struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Enchant        int           `json:"enchant"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Tier             string        `json:"tier"`
				Gems             []interface{} `json:"gems"`
				Enchants         []int         `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"legs"`
			Feet struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Enchant        int           `json:"enchant"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Gems             []interface{} `json:"gems"`
				Enchants         []int         `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"feet"`
			Finger1 struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Enchant        int           `json:"enchant"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Gems             []int         `json:"gems"`
				Enchants         []int         `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"finger1"`
			Finger2 struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Enchant        int           `json:"enchant"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Gems             []int         `json:"gems"`
				Enchants         []int         `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"finger2"`
			Trinket1 struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Gems             []interface{} `json:"gems"`
				Enchants         []interface{} `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"trinket1"`
			Trinket2 struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Gems             []interface{} `json:"gems"`
				Enchants         []interface{} `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"trinket2"`
			Mainhand struct {
				ItemID         int           `json:"item_id"`
				ItemLevel      int           `json:"item_level"`
				Enchant        int           `json:"enchant"`
				Icon           string        `json:"icon"`
				Name           string        `json:"name"`
				ItemQuality    int           `json:"item_quality"`
				IsLegendary    bool          `json:"is_legendary"`
				IsAzeriteArmor bool          `json:"is_azerite_armor"`
				AzeritePowers  []interface{} `json:"azerite_powers"`
				Corruption     struct {
					Added    int `json:"added"`
					Resisted int `json:"resisted"`
					Total    int `json:"total"`
				} `json:"corruption"`
				DominationShards []interface{} `json:"domination_shards"`
				Gems             []interface{} `json:"gems"`
				Enchants         []int         `json:"enchants"`
				Bonuses          []int         `json:"bonuses"`
			} `json:"mainhand"`
		} `json:"items"`
	} `json:"gear"`
}

func ScrapeCharacterData(characterName, realm, region string) (APIResponse, error) {
	//url := fmt.Sprintf("https://raider.io/api/v1/characters/profile?region=%s&realm=%s&name=%s&fields=gear%%2Cmythic_plus_scores_by_season:current",
	url := fmt.Sprintf("https://raider.io/api/v1/characters/profile?region=%s&realm=%s&name=%s&fields=gear%%2Cmythic_plus_scores_by_season:current",
		region, realm, characterName,
	)

	fmt.Println("Fetching data from:", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return APIResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Recieved non-200 status code:", resp.Status)
		return APIResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return APIResponse{}, err
	}

	var data APIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return APIResponse{}, err
	}

	// if len(data.MythicPlusScoresBySeason) > 0 {
	// 	fmt.Println("Score:", data.MythicPlusScoresBySeason[0].Scores.DPS)
	// } else {
	// 	fmt.Println("No Mythic+ score found")
	// }
	return data, err

}
