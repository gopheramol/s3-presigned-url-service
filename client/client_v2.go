package client

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gopheramol/s3-presigned-url-service/configuration"
	"github.com/gopheramol/s3-presigned-url-service/model"
)

type S3Client interface {
	GeneratePresignedURL(ctx context.Context, req model.PreSignedURLRequest) (URL string, err error)
}

type s3Client struct {
	config configuration.HandlerServiceConfig
}

func NewS3Client(
	config configuration.HandlerServiceConfig,
) s3Client {
	return s3Client{config: config}
}

// GeneratePresignedURL generates a presigned URL for accessing an S3 object.
func (client s3Client) GeneratePresignedURL(ctx context.Context, req model.PreSignedURLRequest) (URL string, err error) {

	regionName := client.config.AWSRegion
	secretKey := client.config.SecretKey
	accessKey := client.config.AccessKey

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithRegion(regionName),
	)
	if err != nil {
		log.Fatalf("not able to load aws config %+v", err)
		return
	}

	URL, err = putPresignURL(cfg, req)
	if err != nil {
		log.Fatalf("not able get pre signed url: %+v", err)
	}
	return
}

func putPresignURL(cfg aws.Config, req model.PreSignedURLRequest) (url string, err error) {
	s3client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3client)
	presignedUrl, err := presignClient.PresignPutObject(context.Background(),
		&s3.PutObjectInput{
			Bucket: aws.String(req.BucketName),
			Key:    aws.String(req.File),
		},
		s3.WithPresignExpires(time.Minute*15)) // take it from app config
	if err != nil {
		log.Fatal(err)
		return
	}
	url = presignedUrl.URL
	return

}
