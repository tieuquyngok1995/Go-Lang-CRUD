package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Before request
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		// After request
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		log.Printf("[%s] %s %s | %d | %v | %s",
			method, path, clientIP,
			statusCode, latency,
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)
	}
}
