package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"main/server/api/globals"
	"net/http"
)

func Auth(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user == nil {
		log.Println("User not logged in")
		c.JSON(http.StatusBadRequest, gin.H{"content": "Please login first"})
		c.Abort()
		return
	}
	c.Next()
}
