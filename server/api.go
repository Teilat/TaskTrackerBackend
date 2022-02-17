package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "healthy")
	})

	err := router.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
