package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type CharacterConfig struct {
	CharacterName string `json:"character_name"`
	Realm         string `json:"realm"`
	Region        string `json:"region"`
}

func LoadCharacterConfig(filename string) (CharacterConfig, error) {
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

type GCSConfig struct {
	BucketName string
}

func LoadGCSConfig() (GCSConfig, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		return GCSConfig{}, err
	}
	gcs := &GCSConfig{
		BucketName: os.Getenv("BUCKETNAME"),
	}

	return *gcs, nil

}
