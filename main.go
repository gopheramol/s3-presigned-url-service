package main

import (
	"github.com/gopheramol/s3-presigned-url-service/bootstrap"
	"github.com/gopheramol/s3-presigned-url-service/configuration"
)

func main() {

	config := configuration.LoadConfigs()
	router := bootstrap.InitializeRoutes(config)
	router.Run(":" + config.Port)
}
