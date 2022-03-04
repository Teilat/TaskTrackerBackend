package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	v1 "main/internal/api/v1"
	_ "main/internal/docs"
)

func Init() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		ver1 := api.Group("/v1")
		{
			ver1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
			ver1.GET("/health", v1.HealthCheck)

			tags := ver1.Group("/tag")
			{
				tags.GET("/", v1.GetAllTags)
			}
		}
	}

	address := fmt.Sprintf("%s:%d", viper.Get("api.address"), viper.Get("api.port"))
	log.Printf("Working at %s", address)

	err := router.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}
