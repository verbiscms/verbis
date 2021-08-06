// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/verbiscms/verbis/api/errors"
	"time"
)

// Middleware is the handler function for logging system
// application messages, if there was an error or
// message set, it will be retrieved from the
// context. Status codes between 200 and 400
// will be logged as info, otherwise an
// error.
func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start time
		startTime := time.Now()

		// Set request time for execution
		ctx.Set("request_time", startTime)

		// Processing request
		ctx.Next()

		// Error
		var verbisError *errors.Error
		err, ok := ctx.Get("verbis_error")
		if ok {
			e, ok := err.(*errors.Error)
			if ok {
				verbisError = e
			}
		}

		// Message
		var verbisMessage string
		m, ok := ctx.Get("verbis_message")
		if ok {
			m, ok := m.(string)
			if ok {
				verbisMessage = m
			}
		}

		// End time
		endTime := time.Now()

		// Log fields
		status := ctx.Writer.Status()
		fields := logrus.Fields{
			"status_code":    status,
			"latency_time":   endTime.Sub(startTime),
			"client_ip":      ctx.ClientIP(),
			"request_method": ctx.Request.Method,
			"request_url":    ctx.Request.RequestURI,
			"message":        verbisMessage,
			"error":          verbisError,
		}

		//if strings.Contains(ctx.Request.URL.String(), "/admin") && api.Production {
		//	return
		//}

		if status >= 200 && status < 400 {
			logger.WithFields(fields).Info()
		} else {
			logger.WithFields(fields).Error()
		}
	}
}
