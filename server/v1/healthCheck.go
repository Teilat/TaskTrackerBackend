package v1Api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheck godoc
// @Summary		get health check
// @Tags        accounts
// @Accept      json
// @Produce     json
// @Success     200  {string}  healthy
// @Failure     400  {object}  httputil.HTTPError
// @Failure     404  {object}  httputil.HTTPError
// @Failure     500  {object}  httputil.HTTPError
// @Router      /health [get]
func HealthCheck(g *gin.Context) {
	g.String(http.StatusOK, "healthy")

}
