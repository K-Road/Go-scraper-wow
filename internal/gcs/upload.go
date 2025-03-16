package gcs

import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
)

func UploadtoGCS(bucketName, fileName, objectName string, deleteAfterUpload bool) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	fmt.Println(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Uploading
	object := client.Bucket(bucketName).Object(objectName)
	writer := object.NewWriter(ctx)
	if objectName == "styles.css" {
		writer.ContentType = "text/css"
	}

	_, err = io.Copy(writer, file)
	if err != nil {
		writer.Close()
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	if deleteAfterUpload {
		if fileName != "static/styles.css" && fileName != "styles.css" {
			fmt.Println("Deleting file:", fileName)
			if err := os.Remove(fileName); err != nil && !os.IsNotExist(err) {
				return err
			}
		} else {
			fmt.Println("Skipping deletion for:", fileName)
		}

		if objectName != "styles.css" {
			fmt.Println("Deleting object:", objectName)
			if err := os.Remove(objectName); err != nil && !os.IsNotExist(err) {
				return err
			}
		} else {
			fmt.Println("Skipping deletion for:", objectName)
		}
	}

	fmt.Println("File uploaded")
	return nil
}

func UploadCSStoGCS(bucketName, fileName, objectName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	// bucket := client.Bucket(bucketName)
	// obj := bucket.Object(objectName)
	return nil
}
