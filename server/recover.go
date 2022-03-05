package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/harmony-development/legato/util/panic_fmt"
)

func Recover(clean bool) fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			panic_fmt.RecoverError(clean, e)
			c.Status(500).SendString(fmt.Sprintf("%v", e))
		},
	})
}
