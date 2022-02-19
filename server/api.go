package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "main/server/v1"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func Init() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(gin.BasicAuth()) map[string]string with login/pass

	ver1 := router.Group("/api/v1")
	{
		ver1.GET("/health", v1.HealthCheck)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
