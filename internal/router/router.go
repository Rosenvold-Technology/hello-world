
package router

import (
	"github.com/gin-gonic/gin"

	"hello-website/internal/handlers"
)

func Register(r *gin.Engine) {
	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", handlers.Index)
}
