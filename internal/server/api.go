package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	_ "main/internal/docs"
	"main/internal/server/v1"
)

func Init() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(gin.BasicAuth()) //map[string]string with login/pass

	ver1 := router.Group("/api/v1")
	{
		ver1.GET("/health", v1Api.HealthCheck)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
