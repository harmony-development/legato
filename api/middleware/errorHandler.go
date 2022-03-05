package middleware

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/harmony-development/legato/adapter/hrpc"
	"github.com/harmony-development/legato/api"
)

func (h *MiddlewareHandler) ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		var herr *api.Error
		if errors.As(err, &herr) {
			c.Status(http.StatusBadRequest)
			resp, err := hrpc.MarshalHRPC(&herr.HError, c.Get("Content-Type"))
			if err != nil {
				return err
			}
			return c.Send(resp)
		} else {
			return err
		}
	}
}
