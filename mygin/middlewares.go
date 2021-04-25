package mygin

import (
	"log"
	"time"
)

func MiddlewareV2() HandleFunc {
	return func(c *Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf(
			"[%d] %s in v for group v2\n",
			c.StatusCode,
			c.Req.RequestURI,
			time.Since(t),
			)
	}
}
