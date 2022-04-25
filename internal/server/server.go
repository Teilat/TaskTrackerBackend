package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	_ "main/internal/docs"
	v1 "main/internal/server/api/v1"
)

func Init() {
	//swag init --parseDependency --parseInternal -g main.go

	address := fmt.Sprintf("%s:%d", viper.Get("api.address"), viper.Get("api.port"))

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		ver1 := api.Group("/v1")
		{
			ver1.GET("/", v1.HealthCheck)

			tag := ver1.Group("/tag")
			{
				tag.GET("", v1.GetAllTags)
				tag.POST("", v1.CreateTag)
				tag.PATCH(":id", v1.UpdateTag)
				tag.DELETE(":id", v1.DeleteTag)
				tag.GET("/task", v1.GetTagsByTask)
			}
			task := ver1.Group("/task")
			{
				task.GET("", v1.GetAllTasks)
				task.POST("", v1.GetAllTasks)
				task.PATCH("", v1.UpdateTask)
				task.DELETE("", v1.DeleteTask)
			}
		}
	}

	err := router.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}
