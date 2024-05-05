package bootstrap

import (
	"github.com/gopheramol/s3-presigned-url-service/client"
	"github.com/gopheramol/s3-presigned-url-service/configuration"
	"github.com/gopheramol/s3-presigned-url-service/service"
)

func InitializeObjects(config configuration.HandlerServiceConfig) service.S3HandlerService {

	s3Client := client.NewS3Client(config)

	s3HandlerService := service.NewS3HandlerService(s3Client)

	return s3HandlerService
}
