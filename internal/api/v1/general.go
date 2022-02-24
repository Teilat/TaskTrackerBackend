package v1

import (
	"github.com/gin-gonic/gin"
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
func HealthCheck(g *gin.Context) {
	g.String(http.StatusOK, "OK v1")
}
