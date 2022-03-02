package v1

import (
	"github.com/gin-gonic/gin"
	"main/internal/database/sql"
	"net/http"
)

// HealthCheck  godoc
// @Summary		Health check
// @Tags        General
// @Accept      json
// @Produce     json
// @Success     200  {string}  healthy
// @Router      /health [get]
// @BasePath 	/api/v1
//func RemoveTag(g *gin.Context) {
//	uuid1, _ := uuid.Parse("642878FD-74CC-4C53-B0DC-2C97A1111D10")
//	err := db.RemoveTag(uuid1)
//	if err != nil {
//		log.Fatal(err)
//	}
//	g.String(http.StatusOK, "OK v1")
//}

// GetAllTags  godoc
// @Summary		Get all tags
// @Tags        Tag
// @Accept      json
// @Produce     json
// @Success     200 {json}
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
