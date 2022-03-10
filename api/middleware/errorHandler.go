package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/harmony-development/legato/adapter/hrpc"
	"github.com/harmony-development/legato/api"
)

func (h *MiddlewareHandler) ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		if err == nil {
			return nil
		}

		var herr *api.Error
		if !errors.As(err, &herr) {
			fmt.Println(err) // TODO: Log error to customizeable output
			herr = api.ErrorInternalServerError
		}

		if herr == api.ErrorInternalServerError {
			c.Status(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusBadRequest)
		}

		resp, err := hrpc.MarshalHRPC(&herr.HError, c.Get("Content-Type"))
		if err != nil {
			return err
		}
		return c.Send(resp)
	}
}
