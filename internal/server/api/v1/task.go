package v1

import (
	"github.com/gin-gonic/gin"
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
func GetAllTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := sql.GetDb()
		tags, err := db.GetAllTasks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, tags)
	}
}

// CreateTask  godoc
// @Summary     Create tag
// @Tags        Task
// @Accept      json
// @Produce     json
// @Param       tag body models.AddTask true "add task"
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /task/ [post]
func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.AddTask
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		db := sql.GetDb()
		err = db.CreateNewTask(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, "")
	}
}

// DeleteTask  godoc
// @Summary     Delete task with provided id
// @Tags        Task
// @Accept      json
// @Produce     json
// @Param       id query uuid.UUID true "task id"
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /task/ [delete]
func DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.DeleteTask
		err := c.BindQuery(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		db := sql.GetDb()
		err = db.DeleteTask(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, "")
	}
}

// UpdateTask  godoc
// @Summary     update tag with provided id
// @Tags        Task
// @Accept      json
// @Produce     json
// @Param       task body models.UpdateTask true "update task"
// @Success     200 {object} models.Task
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /task/ [patch]
func UpdateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.UpdateTask
		err := c.BindQuery(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		db := sql.GetDb()
		err = db.UpdateTask(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, "")
	}
}
