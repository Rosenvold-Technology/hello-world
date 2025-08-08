
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		status := c.Writer.Status()
		logger.WithFields(logrus.Fields{
			"status":   status,
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"latency":  latency,
			"clientIP": c.ClientIP(),
		}).Info("request completed")
	}
}
