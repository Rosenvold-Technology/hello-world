package router

import (
	"github.com/gin-gonic/gin"

	"github.com/rosenvold-technologies/hello-world/internal/handlers"
)

func Register(r *gin.Engine) {
	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", handlers.Index)
}
