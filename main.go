package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/database/sql"
)

func main() {
	sql.Connect()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": ""})
	})

	err := router.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = fmt.Scanln()
	if err != nil {
		return
	}
}
