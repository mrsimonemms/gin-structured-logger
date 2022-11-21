# Gin Structured Logger

Structured logger middleware for Gin

## Why?

Shortcuts.

I like using structured logging for my projects and [Zerolog](https://github.com/rs/zerolog)
works nicely. It's not difficult to set up, but it's a bit repetitive, so this is a reusable
library.

If [RequestID](https://github.com/gin-contrib/requestid) is enabled in the application, it
adds a `requestID` to the log output.

Finally, it sets a request-specific instance of the logger that is stored in the Gin context.
There is a `Get` method which can be used to simplify retrieval of the logger from the context.

## Example

> See the [Simple example](./examples//simple) for a runnable example

```go
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
		log := logger.Get(c)

		log.Info().Msg("Hello world endpoint called")

		c.String(200, "hello world")
	})

	r.Run(":3000")
}
```

## Thanks

This is based upon this [Learning Go](https://learninggolang.com/it5-gin-structured-logging.html)
article
