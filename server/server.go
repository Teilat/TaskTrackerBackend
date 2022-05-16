package server

import (
	"fmt"
	session "github.com/ScottHuangZL/gin-jwt-session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"main/server/api/globals"
	"main/server/api/middleware"
	v1 "main/server/api/v1"
	_ "main/server/docs"
)

// @host      localhost:8080
// @BasePath  /api/v1
// @Title     Application Api
// @Version   1.0

func Init() {
	//swag init --parseDependency --parseInternal -g server.go

	address := fmt.Sprintf("%s:%d", viper.Get("api.address"), viper.Get("api.port"))

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	session.NewStore()
	router.Use(session.ClearMiddleware()) //important to avoid mem leak
	router.Use(sessions.Sessions("AuthToken", cookie.NewStore(globals.Secret)))

	router.LoadHTMLGlob("./server/api/templates/*.html")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		apiV1 := api.Group("/v1")
		{
			apiV1.GET("/", v1.HealthCheck())
			tag := apiV1.Group("/tag")
			{
				tag.Use(middleware.Auth)
				tag.GET("", v1.GetAllTags())
				tag.POST("", v1.CreateTag())
				tag.PATCH(":id", v1.UpdateTag())
				tag.DELETE(":id", v1.DeleteTag())
				tag.GET("/task", v1.GetTagsByTask())
			}
			task := apiV1.Group("/task")
			{
				task.Use(middleware.Auth)
				task.GET("", v1.GetAllTasks())
				task.POST("", v1.CreateTask())
				task.PATCH("", v1.UpdateTask())
				task.DELETE("", v1.DeleteTask())
			}
			proj := apiV1.Group("/project")
			{
				proj.Use(middleware.Auth)
				proj.GET("", v1.GetAllProjects())
				proj.POST("", v1.CreateProject())
				proj.PATCH("", v1.UpdateProject())
				proj.DELETE("", v1.DeleteProject())
			}
			authPublic := apiV1.Group("/auth")
			{
				authPublic.GET("/login", v1.LoginGetHandler())
				authPublic.POST("/login", v1.LoginPostHandler())
				authPublic.GET("/", v1.IndexGetHandler())
				authPublic.POST("/register", v1.RegisterHandler())
			}
			authPrivate := apiV1.Group("/auth")
			{
				authPrivate.Use(middleware.Auth)
				authPrivate.GET("/dashboard", v1.DashboardGetHandler())
				authPrivate.GET("/logout", v1.LogoutGetHandler())
			}
		}
	}

	err := router.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}
