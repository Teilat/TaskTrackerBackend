package v1

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"main/internal/server/api/globals"
	"main/internal/server/api/helpers"
	"main/internal/server/api/v1/models"
	"net/http"
)

// LoginGetHandler  godoc
// @Summary     Login get
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Error       400 {string} string
// @Router      /auth/login [get]
func LoginGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.String(http.StatusBadRequest, "Please logout first",
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.String(http.StatusOK, "logged out", gin.H{
			"content": "",
			"user":    user,
		})
	}
}

// LoginPostHandler  godoc
// @Summary     Login post
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       tag body models.Login true "credentials"
// @Success     301
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /auth/login [post]
func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.String(http.StatusBadRequest, "Please logout first", gin.H{"content": "Please logout first"})
			return
		}

		var params models.Login
		err := c.BindJSON(&params)
		if err != nil {
			log.Fatal(err)
			return
		}

		if helpers.EmptyUserPass(params.Username, params.Password) {
			c.String(http.StatusBadRequest, "Parameters can't be empty", gin.H{"content": "Parameters can't be empty"})
			return
		}

		if !helpers.CheckUserPass(params.Username, params.Password) {
			c.String(http.StatusUnauthorized, "Incorrect username or password", gin.H{"content": "Incorrect username or password"})
			return
		}

		session.Set(globals.Userkey, params.Username)
		if err := session.Save(); err != nil {
			c.String(http.StatusInternalServerError, "Failed to save session", gin.H{"content": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/api/v1/auth/dashboard")
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		log.Println("logging out user:", user)
		if user == nil {
			log.Println("Invalid session token")
			return
		}
		session.Delete(globals.Userkey)
		if err := session.Save(); err != nil {
			log.Println("Failed to save session:", err)
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/api/v1/auth/")
	}
}

func DashboardGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		c.String(http.StatusOK, "This is a dashboard", gin.H{
			"content": "This is a dashboard",
			"user":    user,
		})
	}
}