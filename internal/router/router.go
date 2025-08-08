package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()

	// ----- middleware you already had -----
	r.Use(gin.Logger(), gin.Recovery())

	// ----- static files & templates -----
	r.Static("/static", "./web/static")    // css/js
	r.LoadHTMLGlob("web/templates/*.tmpl") // html templates

	// ----- routes -----
	r.GET("/healthz", func(c *gin.Context) { // health probe for CI / K8s
		c.String(http.StatusOK, "ok")
	})

	r.GET("/", func(c *gin.Context) { // landing page
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Title": "Hello, Kubernetes!",
		})
	})

	return r
}
