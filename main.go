package main

import (
	"os"

	"github.com/PolarBearAndrew/go-s3-viewer/server"
)

func main() {

	// Setup env AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION, AWS_BUCKET

	port := ":8080"
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET")

	serv := server.NewS3ViewerServer(server.S3ViewerServConf{
		Port:   port,
		Bucket: bucket,
		Region: region,
	})

	serv.Listen()
}
