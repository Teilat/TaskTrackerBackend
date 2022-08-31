package v1

import (
	"github.com/gin-gonic/gin"
	"main/internal/cache"
	"main/server/api/v1/models"
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
// @Router      /project [get]
func GetAllProjects() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.GetProject
		err := c.BindQuery(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		if params.Id != 0 {
			proj, err := cache.Cache.GetProject(params)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
			} else {
				c.JSON(http.StatusOK, proj)
			}
		} else {
			tags, err := cache.Cache.GetAllProjects()
			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
			} else {
				c.JSON(http.StatusOK, tags)
			}
		}
	}
}

// GetAllTasksByProject  godoc
// @Summary     Get all tasks by project id
// @Tags        Project
// @Accept      json
// @Produce     json
// @Param       projectId query    integer true "projectId"
// @Success     200       {object} models.Task
// @Error       500       {string} string
// @Error       404       {string} string
// @Router      /project/task [get]
func GetAllTasksByProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.TasksByProject
		err := c.BindQuery(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		tags := cache.Cache.GetTasksByProject(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, tags)
		}
	}
}

// GetProjectsTree  godoc
// @Summary     Get tree
// @Tags        Project
// @Accept      json
// @Produce     json
// @Success     200       {array}  models.TreeNode
// @Error       500       {string} string
// @Error       404       {string} string
// @Router      /project/tree [get]
func GetProjectsTree() gin.HandlerFunc {
	return func(c *gin.Context) {
		tags, err := cache.Cache.GetProjectsTree()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, tags)
		}
	}
}

// CreateProject  godoc
// @Summary     Create project
// @Tags        Project
// @Accept      json
// @Produce     json
// @Param       project body     models.AddProject true "add project"
// @Success     200
// @Error       500     {string} string
// @Error       404     {string} string
// @Router      /project [post]
func CreateProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.AddProject
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = cache.Cache.CreateProject(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, "")
		}
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
// @Router      /project [delete]
func DeleteProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.DeleteProject
		err := c.BindQuery(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = cache.Cache.DeleteProject(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, "")
		}
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
// @Router      /project [patch]
func UpdateProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.UpdateProject
		err := c.BindQuery(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = cache.Cache.UpdateProject(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, "")
		}
	}
}
