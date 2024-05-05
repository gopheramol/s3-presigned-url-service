package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/gopheramol/s3-presigned-url-service/configuration"
	"github.com/gopheramol/s3-presigned-url-service/controller"
)

func InitializeRoutes(config configuration.HandlerServiceConfig) *gin.Engine {

	s3HandlerService := InitializeObjects(config)

	router := gin.Default()

	handlerAPI := router.Group("/api/v1")
	{
		handlerAPI.POST("/get-presigned-url", controller.NewController(s3HandlerService.GeneratePresignedURL).HandlerFunc)

	}
	return router
}
