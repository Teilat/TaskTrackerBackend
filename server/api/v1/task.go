package v1

import (
	"github.com/gin-gonic/gin"
	"main/internal/cache"
	"main/server/api/v1/models"
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
// @Router      /task [get]
func GetAllTasks() gin.HandlerFunc {
	return func(c *gin.Context) {
		tags := cache.Cache.GetAllTasks()
		c.JSON(http.StatusOK, tags)
	}
}

// CreateTask  godoc
// @Summary     Create task
// @Tags        Task
// @Accept      json
// @Produce     json
// @Param       tag body models.AddTask true "add task"
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /task [post]
func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.AddTask
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		err = cache.Cache.CreateTask(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, "")
		}
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
// @Router      /task [delete]
func DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.DeleteTask
		err := c.BindQuery(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		err = cache.Cache.DeleteTask(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, "")
		}
	}
}

// UpdateTaskPos  godoc
// @Summary     Update task column
// @Tags        Task
// @Accept      json
// @Produce     json
// @Param       task body models.UpdateTaskPos true "Update column"
// @Error       500       {string} string
// @Error       404       {string} string
// @Router      /task/column [patch]
func UpdateTaskPos() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.UpdateTaskPos
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = cache.Cache.UpdateTaskPos(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, "")
		}
	}
}

// UpdateTask  godoc
// @Summary     update tag with provided id
// @Tags        Task
// @Accept      json
// @Produce     json
// @Param       task body models.UpdateTask true "Update task"
// @Success     200 {object} models.Task
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /task [patch]
func UpdateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.UpdateTask
		err := c.BindQuery(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		err = cache.Cache.UpdateTask(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			c.JSON(http.StatusOK, "")
		}
	}
}
