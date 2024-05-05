package configuration

import (
	"time"
)

type HandlerServiceConfig struct {
	BucketName string
	ObjectKey  string
	AccessKey  string
	SecretKey  string
	AWSRegion  string
	Expiry     time.Duration
}
