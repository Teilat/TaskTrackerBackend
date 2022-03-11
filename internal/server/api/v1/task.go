package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"main/internal/database/sql"
	"main/internal/server/api/v1/models"
	"net/http"
)

// GetAllTasks  godoc
// @Summary     Get all tasks
// @Tags        Task
// @Accept      json
// @Produce     json
// @Success     200 {array}  models.Task
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /task/ [get]
func GetAllTasks(g *gin.Context) {
	db := sql.GetDb()
	tags, err := db.GetAllTasks()
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, tags)
}

// CreateTask  godoc
// @Summary     Create tag
// @Tags        Task
// @Accept      json
// @Produce     json
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /task/ [post]
func CreateTask(g *gin.Context) {

	var params models.AddTask
	err := g.ShouldBindBodyWith(&params, binding.JSON)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}

	db := sql.GetDb()
	err = db.CreateNewTask(params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, "")
}

// DeleteTask  godoc
// @Summary     Delete tag with provided id
// @Tags        Task
// @Accept      json
// @Produce     json
// @Param       id query uuid.UUID true "task id"
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /task/ [delete]
func DeleteTask(g *gin.Context) {

	var params models.DeleteTask
	err := g.BindQuery(&params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}

	db := sql.GetDb()
	err = db.DeleteTask(params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, "")
}

// UpdateTask  godoc
// @Summary     update tag with provided id
// @Tags        Task
// @Accept      json
// @Produce     json
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /task/ [put]
func UpdateTask(g *gin.Context) {

	var params models.UpdateTask
	err := g.BindQuery(&params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}

	db := sql.GetDb()
	err = db.UpdateTask(params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, "")
}
