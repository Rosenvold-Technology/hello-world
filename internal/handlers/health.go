package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthz responds with 200 "ok"
// Kubernetes and CI both use it for liveness + readiness.
func Healthz(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
