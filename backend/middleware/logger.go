package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"strings"
	"sync"
	"time"
)

// Logger returns a logger middleware which run after all requests to get the response data
func Logger() fiber.Handler {
	var (
		once       sync.Once
		errHandler fiber.ErrorHandler
	)

	return func(c *fiber.Ctx) (err error) {
		once.Do(func() {
			errHandler = c.App().ErrorHandler
		})

		var start, stop time.Time
		start = time.Now()
		chainErr := c.Next()

		// Manually call error handler
		if chainErr != nil {
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}
		stop = time.Now()

		log.WithFields(log.Fields{
			"ip":         c.IP(),
			"latency":    strings.TrimSpace(fmt.Sprintf("%7v", stop.Sub(start).Round(time.Millisecond))),
			"path":       c.Path(),
			"userAgent":  c.Get(fiber.HeaderUserAgent),
			"statusCode": c.Response().StatusCode(),
		}).Info()
		return nil
	}
}
