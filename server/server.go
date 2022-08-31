package server

import (
	"fmt"
	session "github.com/ScottHuangZL/gin-jwt-session"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"main/server/api/globals"
	v1 "main/server/api/v1"
	_ "main/server/docs"
	"time"
)

// @host      localhost:8080
// @Title     Application Api
// @Version   1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name token

func Init() {
	//swag init --parseDependency --parseInternal -g server.go
	address := fmt.Sprintf("%s:%d", viper.Get("api.address"), viper.Get("api.port"))

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	session.NewStore()
	cookieStore := cookie.NewStore(globals.Secret)
	cookieStore.Options(sessions.Options{HttpOnly: false, Secure: false, MaxAge: 86400, Path: "/"})
	router.Use(session.ClearMiddleware()) //important to avoid mem leak
	router.Use(sessions.Sessions("token", cookieStore))
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost", "http://localhost:8080", "http://localhost:*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Authorization", "Origin", "Accept", "X-Requested-With", "Content-Type", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/login", v1.LoginPostHandler())
	router.POST("/register", v1.RegisterHandler())
	router.GET("/logout", v1.LogoutGetHandler())

	router.GET("/", v1.HealthCheck())
	tag := router.Group("/tag")
	{
		//router.Use(middleware.Auth)
		tag.GET("", v1.GetAllTags())
		tag.POST("", v1.CreateTag())
		tag.PATCH(":id", v1.UpdateTag())
		tag.DELETE(":id", v1.DeleteTag())
		tag.GET("/task", v1.GetTagsByTask())
	}
	task := router.Group("/task")
	{
		//router.Use(middleware.Auth)
		task.GET("", v1.GetAllTasks())
		task.PATCH("/column", v1.UpdateTaskPos())
		task.POST("", v1.CreateTask())
		task.PATCH("", v1.UpdateTask())
		task.DELETE("", v1.DeleteTask())
	}
	proj := router.Group("/project")
	{
		//router.Use(middleware.Auth)
		proj.GET("", v1.GetAllProjects())
		proj.GET("/task", v1.GetAllTasksByProject())
		proj.GET("/tree", v1.GetProjectsTree())
		proj.POST("", v1.CreateProject())
		proj.PATCH("", v1.UpdateProject())
		proj.DELETE("", v1.DeleteProject())
	}
	err := router.Run(address)
	if err != nil {
		log.Fatal(err)
	}
}
