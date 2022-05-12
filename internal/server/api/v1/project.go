package v1

import (
	"github.com/gin-gonic/gin"
	"main/internal/database/sql"
	"main/internal/server/api/v1/models"
	"net/http"
)

// GetAllProjects  godoc
// @Summary     Get all projects
// @Tags        Project
// @Accept      json
// @Produce     json
// @Success     200 {array}  models.Project
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /project/ [get]
func GetAllProjects() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := sql.GetDb()
		tags, err := db.GetAllProjects()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, tags)
	}
}

// CreateProject  godoc
// @Summary     Create tag
// @Tags        Project
// @Accept      json
// @Produce     json
// @Param       tag body models.AddProject true "add project"
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /project/ [post]
func CreateProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.AddProject
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		db := sql.GetDb()
		err = db.CreateProject(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, "")
	}
}

// DeleteProject  godoc
// @Summary     Delete project with provided id
// @Tags        Project
// @Accept      json
// @Produce     json
// @Param       id query uuid.UUID true "project id"
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /project/ [delete]
func DeleteProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.DeleteProject
		err := c.BindQuery(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		db := sql.GetDb()
		err = db.DeleteProject(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, "")
	}
}

// UpdateProject  godoc
// @Summary     update project with provided id
// @Tags        Project
// @Accept      json
// @Produce     json
// @Param       project body models.UpdateProject true "update project"
// @Success     200 {object} models.Project
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /project/ [patch]
func UpdateProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.UpdateProject
		err := c.BindQuery(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		db := sql.GetDb()
		err = db.UpdateProject(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, "")
	}
}
