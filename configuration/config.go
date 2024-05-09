package configuration

import (
	"log"
	"os"
	"time"

	"github.com/gopheramol/s3-presigned-url-service/util"
	"github.com/joho/godotenv"
)

type HandlerServiceConfig struct {
	BucketName string
	ObjectKey  string
	AccessKey  string
	SecretKey  string
	AWSRegion  string
	Expiry     time.Duration
	Port       string
}

func LoadConfigs() HandlerServiceConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")
	accessKey := os.Getenv("ACCESS_KEY")
	region := os.Getenv("REGION")
	expiry := os.Getenv("EXPIRY")
	port := os.Getenv("PORT")
	expirationDuration := util.GetExpiryTime(expiry)

	config := HandlerServiceConfig{
		BucketName: s3Bucket,
		SecretKey:  secretKey,
		AccessKey:  accessKey,
		AWSRegion:  region,
		Expiry:     expirationDuration,
		Port:       port,
	}
	return config
}
