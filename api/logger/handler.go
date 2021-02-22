// Copyright 2020 The Verbis Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

// Log
//
//
func Handler() gin.HandlerFunc {
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
		fields := log.Fields{
			"status_code":    status,
			"latency_time":   endTime.Sub(startTime),
			"client_ip":      ctx.ClientIP(),
			"request_method": ctx.Request.Method,
			"request_url":    ctx.Request.RequestURI,
			"message":        verbisMessage,
			"error":          verbisError,
		}

		if status >= 200 && status < 400 {
			log.WithFields(fields).Info()
		} else {
			log.WithFields(fields).Error()
		}
	}
}
