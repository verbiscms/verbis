package middleware

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func Log() gin.HandlerFunc {
	return func(g *gin.Context) {
		// Start time
		startTime := time.Now()
		// Set request time for execution
		g.Set("request_time", startTime)
		// Processing request
		g.Next()
		// Error
		var verbisError errors.Error
		e, exists := g.Get("verbis_error"); if exists {
			e, ok := e.(*errors.Error); if ok {
				verbisError = *e
			}
		}
		// Message
		var verbisMessage string
		m, exists := g.Get("verbis_message"); if exists {
			m, ok := m.(string); if ok {
				verbisMessage = m
			}
		} else if verbisError.Message != "" {
			verbisMessage = verbisError.Message
		} else {
			verbisMessage = ""
		}
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
		// Data Length
		dataLength := g.Writer.Size()
		// User agent
		clientUserAgent := g.Request.UserAgent()
		// Log fields
		fields := log.Fields{
			"status_code"  		: statusCode,
			"latency_time" 		: latencyTime,
			"client_ip"    		: clientIP,
			"request_method"   	: reqMethod,
			"request_url"      	: reqUri,
			"data_length"   	: dataLength,
			"user_agent"    	: clientUserAgent,
			"message"			: verbisMessage,
			"error"				: verbisError,
		}
		// Log format
		if verbisError.Code == errors.TEMPLATE {
			fields["status_code"] = 500
			log.WithFields(fields).Error()
		} else if statusCode >= 200 && statusCode < 400 {
			log.WithFields(fields).Info()
		} else {
			log.WithFields(fields).Error()
		}
	}
}


