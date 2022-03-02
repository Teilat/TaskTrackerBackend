package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "main/internal/api/v1"
	_ "main/internal/docs"
)

func Init() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	ver1 := router.Group("/api/v1")
	{
		ver1.GET("/health", v1.HealthCheck)

		tags := ver1.Group("/tag")
		{
			tags.GET("/", v1.GetAllTags)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
