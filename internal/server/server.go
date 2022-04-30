package server

import (
	"fmt"
	"github.com/ScottHuangZL/gin-jwt-session"
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

	//below are optional setting, you change it or just comment it to let it as default
	// session.SecretKey = "You any very secriet key !@#$!@%@"  //Any characters
	// session.JwtTokenName = "YouCanChangeTokenName"               //no blank character
	// session.DefaultFlashSessionName = "YouCanChangeTheFlashName" //no blank character
	// session.DefaultSessionName = "YouCanChangeTheSessionName"    //no blank character
	//end of optional setting

	session.NewStore()
	router.Use(session.ClearMiddleware()) //important to avoid mem leak

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setupRouter(router)

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
				task.POST("", v1.CreateTask)
				task.PATCH("", v1.UpdateTask)
				task.DELETE("", v1.DeleteTask)
			}
			proj := ver1.Group("/project")
			{
				proj.GET("", v1.GetAllProjects)
				proj.POST("", v1.CreateProject)
				proj.PATCH("", v1.UpdateProject)
				proj.DELETE("", v1.DeleteProject)
			}
		}
	}

	err := router.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter(r *gin.Engine) {
	r.Delims("{%", "%}")
	// Default With the Logger and Recovery middleware already attached
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.GET("/login", v1.LoginHandler)
	r.GET("/logout", v1.LoginHandler) //logout also leverage login handler, since it just need clear session
	r.POST("/validate-jwt-login", v1.ValidateJwtLoginHandler)
	r.GET("/index.html", v1.HomeHandler)
	r.GET("/index", v1.HomeHandler)
	r.GET("", v1.HomeHandler)

	r.GET("/some-cookie-example", v1.SomeCookiesHandler)

}
