package v1

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"main/db/sql"
	"main/server/api/globals"
	"main/server/api/helpers"
	"main/server/api/v1/models"
	"net/http"
)

// LoginGetHandler  godoc
// @Summary     go to login page
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Error       400 {string} string
// @Router      /login [get]
func LoginGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"content": "",
			"user":    user,
		})
	}
}

// LoginPostHandler  godoc
// @Summary     Login user
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       tag body models.Login true "credentials"
// @Success     301
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /login [post]
func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.JSON(http.StatusBadRequest, gin.H{"content": "Please logout first"})
			return
		}
		var params models.Login

		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		if helpers.EmptyUserPass(params) {
			c.JSON(http.StatusBadRequest, gin.H{"content": "Parameters can't be empty"})
			return
		}

		if !helpers.CheckUserPass(params) {
			c.JSON(http.StatusUnauthorized, gin.H{"content": "Incorrect username or password"})
			return
		}

		session.Set(globals.Userkey, params.Username)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to save session"})
			return
		}
	}
}

// LogoutGetHandler  godoc
// @Summary     Logout user
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Success     301
// @Error       500 {string} string
// @Router      /logout [get]
func LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		log.Println("logging out user:", user)
		if user == nil {
			log.Println("Invalid session token")
		}
		session.Delete(globals.Userkey)
		if err := session.Save(); err != nil {
			log.Println("Failed to save session:", err)
		}
	}
}

// RegisterHandler  godoc
// @Summary     register user
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       tag body models.AddUser true "user"
// @Success     200
// @Error       500 {string} string
// @Router      /register [post]
func RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.AddUser
		err := c.BindJSON(&params)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to register"})
		}

		db := sql.GetDb()
		err = db.CreateNewUser(params)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to register"})
		}
	}
}

func IndexGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		c.JSON(http.StatusOK, gin.H{
			"content": "This is an index page...",
			"user":    user,
		})
	}
}

func DashboardGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		c.JSON(http.StatusOK, gin.H{
			"content": "This is a dashboard",
			"user":    user,
		})
	}
}
