package middleware

import (
	"errors"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/harmony-development/legato/api"
)

func (h *MiddlewareHandler) Auth(path string) fiber.Handler {
	enableAuth := api.MethodMetadata[path].RequiresAuthentication
	return func(c *fiber.Ctx) error {
		if !enableAuth {
			return c.Next()
		}
		session := c.Get("Authorization")
		if session == "" {
			protocols := strings.Split(c.Get("Sec-WebSocket-Protocol"), ",")
			session = protocols[len(protocols)-1]
		}
		if session == "" {
			return api.ErrorInvalidAuth
		}

		userID, err := h.db.GetUserForSession(c.Context(), session)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				return api.ErrorInvalidAuth
			}
			return err
		}

		ctx, _ := c.UserContext().(api.LegatoContext)
		ctx.UserID = userID

		// nolint
		return c.Next()
	}
}
