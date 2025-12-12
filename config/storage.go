package config

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

var storageClient *storage.Client
var bucketName string

func ConnectStorage() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	storageClient = client
	bucketName = os.Getenv("GCS_BUCKET_NAME")

	defer client.Close()
	log.Println("Storage connected successfully")
}

func UploadFile(ctx context.Context, fileName string, file io.Reader) error {
	// 設定逾時
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// 建立 writer 並上傳
	writer := storageClient.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	if _, err := io.Copy(writer, file); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	log.Printf("檔案已上傳到 bucket %v 的 %v\n", bucketName, fileName)

	return nil
}
