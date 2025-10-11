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
		path := c.FullPath()

		if status >= http.StatusBadRequest {

			if len(c.Errors) > 0 {
				logEvent := elog.Error().
					Int("status", status).
					Str("method", c.Request.Method).
					Str("path", path)

				for _, err := range c.Errors {
					logEvent = logEvent.Err(err.Err)
				}

				if status >= http.StatusInternalServerError {
					logEvent.Msg("SERVER ERROR")
				} else {
					logEvent.Str("client_ip", c.ClientIP()).
						Msg("CLIENT ERROR - Bad request/validation failed")
				}

			} else {
				elog.Warn().
					Int("status", status).
					Str("path", path).
					Msg("HTTP status error detected, but no Gin error attached.")
			}

		} else if status >= http.StatusOK && status < http.StatusBadRequest {
			rlog.Info().
				Str("path", path).
				Int("status", status).
				Msg("endpoint worked")
		}
	}
}
