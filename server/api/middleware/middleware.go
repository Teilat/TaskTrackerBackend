package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Auth(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		log.Println("User not logged in")
		c.Redirect(http.StatusMovedPermanently, "/api/v1/auth/login")
		c.Abort()
		return
	}
	c.Next()
}

func Header(c *gin.Context) {
	c.Request.Header.Set("Access-Control-Allow-Origin", "*")
	c.Next()
}
