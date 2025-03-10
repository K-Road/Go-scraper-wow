package gcs

import (
	"context"
	"fmt"
	"io"
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

	_, err = io.Copy(writer, file)
	if err != nil {
		writer.Close()
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	if err := os.Remove(fileName); err != nil {
		return err
	}
	if fileName != "index.html" {
		if err := os.Remove("index.html"); err != nil && !os.IsNotExist(err) {
			return err
		}
	}

	fmt.Println("File uploaded")
	return nil

}
