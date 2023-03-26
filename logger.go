package ginstructuredlogger

import (
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	CtxKey = "logger"
)

func Get(c *gin.Context) *zerolog.Logger {
	return c.MustGet(CtxKey).(*zerolog.Logger)
}

func New() gin.HandlerFunc {
	return func(c *gin.Context) {
		logCtx := (&log.Logger).With()

		// Start timer
		start := time.Now()

		if requestId := requestid.Get(c); requestId != "" {
			logCtx = logCtx.Str("requestID", requestId)
		}
		logger := logCtx.Logger()

		// Store the logger so it can be used in contexts
		c.Set(CtxKey, &logger)

		// Continue to the next middleware
		c.Next()

		// Stop timer and calculate the latency
		timeStamp := time.Now()
		latency := timeStamp.Sub(start)
		if latency > time.Minute {
			latency = latency.Truncate(time.Second)
		}

		raw := c.Request.URL.RawQuery
		path := c.Request.URL.Path
		if raw != "" {
			path = path + "?" + raw
		}

		var logEvent *zerolog.Event
		if c.Writer.Status() >= 500 {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Debug()
		}

		logEvent.
			Str("clientIP", c.ClientIP()).
			Str("method", c.Request.Method).
			Int("statusCode", c.Writer.Status()).
			Int("bodySize", c.Writer.Size()).
			Str("path", path).
			Str("latency", latency.String()).
			Msg(c.Errors.ByType(gin.ErrorTypePrivate).String())
	}
}
