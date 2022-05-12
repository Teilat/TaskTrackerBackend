package v1

import (
	"github.com/gin-gonic/gin"
	"main/database/sql"
	"main/server/api/v1/models"
	"net/http"
)

// GetAllTags  godoc
// @Summary     Get all tags
// @Tags        Tag
// @Accept      json
// @Produce     json
// @Success     200 {array}  models.Tag
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /tag/ [get]
func GetAllTags() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := sql.GetDb()
		tags, err := db.GetAllTags()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, tags)
	}
}

// CreateTag  godoc
// @Summary     Create tag
// @Tags        Tag
// @Accept      json
// @Produce     json
// @Param       tag body models.AddTag true "add tag"
// @Success     200 {object} models.Tag
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /tag/ [post]
func CreateTag() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.AddTag
		err := c.BindQuery(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		db := sql.GetDb()
		err = db.CreateNewTag(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, "")
	}
}

// DeleteTag  godoc
// @Summary     Delete tag
// @Tags        Tag
// @Accept      json
// @Produce     json
// @Param       id path string true "tag id"
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /tag/{id} [delete]
func DeleteTag() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.DeleteTag
		err := c.BindQuery(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		db := sql.GetDb()
		err = db.DeleteTag(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, "")
	}
}

// UpdateTag  godoc
// @Summary     Update tag
// @Tags        Tag
// @Accept      json
// @Produce     json
// @Param       tag body models.UpdateTag true "update tag"
// @Success     200 {object} models.Tag
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /tag/{id} [patch]
func UpdateTag() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.UpdateTag
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		db := sql.GetDb()
		err = db.UpdateTag(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, "")
	}
}

// GetTagsByTask  godoc
// @Summary     Get tags by task
// @Tags        Tag
// @Accept      json
// @Produce     json
// @Param       id path string true "Task ID"
// @Success     200 {array}  models.Tag
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /tag/{id} [get]
func GetTagsByTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params models.TagsByTask
		err := c.BindQuery(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		db := sql.GetDb()
		data, err := db.GetTagsByTask(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, data)
	}
}
