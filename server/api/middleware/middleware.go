package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"main/server/api/globals"
)

func Auth(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user == nil {
		log.Println("User not logged in")
		//c.Redirect(http.StatusMovedPermanently, "/login")
		//c.Abort()
		//return
	}
	c.Next()
}
