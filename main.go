package main

import (
	"log"
	"os"

	"github.com/gopheramol/s3-presigned-url-service/bootstrap"
	"github.com/gopheramol/s3-presigned-url-service/configuration"
	"github.com/gopheramol/s3-presigned-url-service/util"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")
	accessKey := os.Getenv("ACCESS_KEY")
	region := os.Getenv("REGION")
	expiry := os.Getenv("EXPIRY")

	expirationDuration := util.GetExpiryTime(expiry)

	config := configuration.HandlerServiceConfig{
		BucketName: s3Bucket,
		SecretKey:  secretKey,
		AccessKey:  accessKey,
		AWSRegion:  region,
		Expiry:     expirationDuration,
	}

	router := bootstrap.InitializeRoutes(config)
	router.Run(":8080")
}
