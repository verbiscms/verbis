package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func LogMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		// Start time
		startTime := time.Now()
		// Processing request
		g.Next()
		// End time
		endTime := time.Now()
		// Execution time
		latencyTime := endTime.Sub(startTime)
		// Request mode
		reqMethod := g.Request.Method
		// Request routing
		reqUri := g.Request.RequestURI
		// Status code
		statusCode := g.Writer.Status()
		// Request IP
		clientIP := g.ClientIP()
		// Log format
		log.WithFields(log.Fields{
			"status_code"  : statusCode,
			"latency_time" : latencyTime,
			"client_ip"    : clientIP,
			"req_method"   : reqMethod,
			"req_uri"      : reqUri,
		}).Info()
	}
}

