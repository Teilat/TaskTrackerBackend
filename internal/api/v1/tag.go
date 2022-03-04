package v1

import (
	"github.com/gin-gonic/gin"
	_ "main/internal/api/v1/models"
	"main/internal/database/sql"
	"net/http"
)

// GetAllTags  godoc
// @Summary		Get all tags
// @Tags        Tag
// @Accept      json
// @Produce     json
// @Success     200 {array} model.Tag
// @Error		500	{json}
// @Router      / [get]
// @BasePath 	/api/v1/tag
func GetAllTags(g *gin.Context) {
	db := sql.GetDb()
	tags, err := db.GetAllTags()
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
	}
	g.JSON(http.StatusOK, tags)
}
