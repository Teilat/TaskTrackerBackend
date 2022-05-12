package server

import (
	"fmt"
	"github.com/ScottHuangZL/gin-jwt-session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	_ "main/internal/docs"
	"main/internal/server/api/globals"
	"main/internal/server/api/middleware"
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
	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

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
			auth := ver1.Group("/auth")
			{
				auth.GET("/login", v1.LoginGetHandler())
				auth.POST("/login", v1.LoginPostHandler())
				auth.GET("/dashboard", v1.DashboardGetHandler())
				auth.GET("/logout", v1.LogoutGetHandler())
				router.Use(middleware.Auth)
			}
		}
	}

	err := router.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}
