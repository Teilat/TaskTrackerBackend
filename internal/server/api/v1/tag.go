package v1

import (
	"github.com/gin-gonic/gin"
	"main/internal/database/sql"
	"main/internal/server/api/v1/models"
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
func GetAllTags(g *gin.Context) {
	db := sql.GetDb()
	tags, err := db.GetAllTags()
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, tags)
}

// CreateTag  godoc
// @Summary     Create tag
// @Tags        Tag
// @Accept      json
// @Produce     json
// @Param       name query string true "tag name"
// @Param       color query string true "tag color" minlength(6) maxlength(6)
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /tag/ [post]
func CreateTag(g *gin.Context) {

	var params models.AddTag
	err := g.BindQuery(&params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}

	db := sql.GetDb()
	err = db.CreateNewTag(params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, "")
}

// DeleteTag  godoc
// @Summary     Delete tag with provided id
// @Tags        Tag
// @Accept      json
// @Produce     json
// @Param       id query uuid.UUID true "tag id"
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /tag/ [delete]
func DeleteTag(g *gin.Context) {

	var params models.DeleteTag
	err := g.BindQuery(&params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}

	db := sql.GetDb()
	err = db.DeleteTag(params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, "")
}

// UpdateTag  godoc
// @Summary     update tag with provided id
// @Tags        Tag
// @Accept      json
// @Produce     json
// @Param       id query uuid.UUID true "tag id"
// @Param       name query string false "tag name"
// @Param       color query string false "tag color"
// @Success     200
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /tag/ [put]
func UpdateTag(g *gin.Context) {

	var params models.UpdateTag
	err := g.BindQuery(&params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}

	db := sql.GetDb()
	err = db.UpdateTag(params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, "")
}

// GetTagsByTask godoc
// @Summary      get tags which contains task by provided id
// @Tags         Tag
// @Accept       json
// @Produce      json
// @Param        id query uuid.UUID true "tag id"
// @Success      200 {array} models.Tag
// @Error        500 {string} string
// @Error        404 {string} string
// @Router       /tag/task [get]
func GetTagsByTask(g *gin.Context) {

	var params models.TagsByTask
	err := g.BindQuery(&params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}

	db := sql.GetDb()
	data, err := db.GetTagsByTask(params)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}

	g.JSON(http.StatusOK, data)
}
