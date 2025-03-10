package gcs

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/storage"
)

func UploadtoGCS(bucketName, fileName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Uploading
	object := client.Bucket(bucketName).Object("index.html")
	writer := object.NewWriter(ctx)
	defer writer.Close()

	_, err = file.WriteTo(writer)
	if err != nil {
		return err
	}

	fmt.Println("File uploaded")
	return nil

}
