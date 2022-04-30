package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
func GetAllProjects(g *gin.Context) {
	db := sql.GetDb()
	tags, err := db.GetAllProjects()
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, tags)
}

// CreateProject  godoc
// @Summary     Create tag
// @Tags        Project
// @Accept      json
// @Produce     json
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /project/ [post]
func CreateProject(g *gin.Context) {

	var params models.AddProject
	err := g.ShouldBindBodyWith(&params, binding.JSON)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}

	db := sql.GetDb()
	err = db.CreateProject(params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, "")
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
func DeleteProject(g *gin.Context) {

	var params models.DeleteProject
	err := g.BindQuery(&params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}

	db := sql.GetDb()
	err = db.DeleteProject(params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, "")
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
func UpdateProject(g *gin.Context) {

	var params models.UpdateProject
	err := g.BindQuery(&params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}

	db := sql.GetDb()
	err = db.UpdateProject(params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, "")
}
