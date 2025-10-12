package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/not-raghu/go-app/utils"
)

func ReqLogger() gin.HandlerFunc {
	rlog, elog := utils.GetLog()

	return func(c *gin.Context) {
		c.Next()

		status := c.Writer.Status()
		method := c.Request.Method
		path := c.FullPath()
		clientIP := c.ClientIP()

		if len(c.Errors) > 0 {
			logEvent := elog.Error().
				Int("status", status).
				Str("method", method).
				Str("path", path).
				Str("client_ip", clientIP)

			for _, err := range c.Errors {
				logEvent = logEvent.Err(err.Err)
			}

			logEvent.Msg("Request failed with errors")
			return
		}

		if status >= http.StatusInternalServerError {
			elog.Error().
				Int("status", status).
				Str("method", method).
				Str("path", path).
				Str("client_ip", clientIP).
				Msg("Server error")
		} else if status >= http.StatusBadRequest {
			elog.Warn().
				Int("status", status).
				Str("method", method).
				Str("path", path).
				Str("client_ip", clientIP).
				Msg("Client error")
		} else {
			rlog.Info().
				Int("status", status).
				Str("method", method).
				Str("path", path).
				Str("client_ip", clientIP).
				Msg("OK")
		}
	}
}
