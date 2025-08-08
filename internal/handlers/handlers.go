
package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "🌟 The World's Most Over‑Engineered Hello Website 🌟",
		"message": "Hello, Kubernetes!",
		"time":    time.Now().Format(time.RFC1123),
	})
}
