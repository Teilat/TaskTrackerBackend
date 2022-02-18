package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce string
// @Success 200 {string} healthy
// @Router /health [get]
func HealthCheck(g *gin.Context) {
	g.String(http.StatusOK, "healthy")

}
