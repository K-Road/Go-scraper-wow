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

	//check version
	if objectName == "styles.css" && !shouldUpload(client, ctx, bucketName, objectName, fileName) {
		fmt.Println("File not newer. Skip upload")
		return nil
	}

	//Uploading
	if err := upload(client, ctx, bucketName, objectName, file, objectName == "styles.css"); err != nil {
		return err
	}

	handleFileDeletion(fileName, objectName, deleteAfterUpload)

	fmt.Println("File uploaded")
	return nil
}

func upload(client *storage.Client, ctx context.Context, bucketName, objectName string, file *os.File, isCSS bool) error {
	object := client.Bucket(bucketName).Object(objectName)
	writer := object.NewWriter(ctx)
	defer writer.Close()

	if isCSS {
		writer.ContentType = "text/css"
		writer.Metadata = map[string]string{
			"Cache-Control": "no-cache",
		}
	}

	if _, err := io.Copy(writer, file); err != nil {
		return fmt.Errorf("copy to GCS: %w", err)
	}

	return nil
}

func handleFileDeletion(fileName, objectName string, deleteAfterUpload bool) {
	if !deleteAfterUpload {
		return
	}

	if fileName != "static/styles.css" && fileName != "styles.css" {
		fmt.Println("Deleting file:", fileName)
		if err := os.Remove(fileName); err != nil && !os.IsNotExist(err) {
			fmt.Println("Error deleting file:", err)
		}
	} else {
		fmt.Println("Skipping deletion for:", fileName)
	}

	if objectName != "styles.css" {
		fmt.Println("Deleting object:", objectName)
		if err := os.Remove(objectName); err != nil && !os.IsNotExist(err) {
			fmt.Println("Error deleting object:", err)
		}
	} else {
		fmt.Println("Skipping deletion for:", objectName)
	}
}

func shouldUpload(client *storage.Client, ctx context.Context, bucketName, objectName, fileName string) bool {
	object := client.Bucket(bucketName).Object(objectName)
	attrs, err := object.Attrs(ctx)

	if err == nil {
		localFileInfo, err := os.Stat(fileName)
		if err != nil {
			return false
		}
		return localFileInfo.ModTime().After(attrs.Updated)
	}
	return err == storage.ErrObjectNotExist
}
