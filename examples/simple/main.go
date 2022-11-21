package main

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"

	logger "github.com/mrsimonemms/gin-structured-logger"
)

func main() {
	r := gin.New()

	// RequestID is optional
	r.Use(requestid.New(), logger.New())

	r.GET("/", func(c *gin.Context) {
		logger.Get(c).Info().Msg("Hello world endpoint called")

		c.String(200, "hello world")
	})

	r.Run(":3000")
}
