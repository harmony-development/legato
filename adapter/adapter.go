// Package adapter is an adapter for the hrpc protocol.
package adapter

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	goserver "github.com/harmony-development/hrpc/pkg/go-server"
	"github.com/harmony-development/legato/adapter/hrpc"
	"github.com/harmony-development/legato/api/middleware"
	"github.com/harmony-development/legato/util/panic_fmt"
)

func RegisterHandler[T context.Context](mid *middleware.MiddlewareHandler, router fiber.Router, handler goserver.MethodHandler[T]) {
	for route, h := range handler.Routes() {
		router.All(route, mid.Auth(route), UnaryHandler(h))
	}
	for route, h := range handler.StreamRoutes() {
		router.All(route, mid.Auth(route), StreamHandler(h))
	}
}

func UnaryHandler[T context.Context](h goserver.UnaryMethodData[T]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		input, err := hrpc.UnmarshalHRPC(c.Body(), c.Get("Content-Type"), h.Input)
		if err != nil {
			return err
		}
		result, err := h.Handler(c.UserContext().(T), input)
		if err != nil {
			return err
		}
		response, err := hrpc.MarshalHRPC(result, c.Get("Content-Type"))
		return c.Send(response)
	}
}

func StreamHandler[T context.Context](h goserver.StreamMethodData[T]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		return websocket.New(func(conn *websocket.Conn) {
			reqChan := make(chan goserver.VTProtoMessage)
			sendChan := make(chan goserver.VTProtoMessage)
			go func() {
				defer func() {
					close(reqChan)
					close(sendChan)

					if r := recover(); r != nil {
						panic_fmt.RecoverError(os.Getenv("RAW_LOG") != "true", r)
					}
				}()

				h.Handler(ctx.(T), reqChan, sendChan)
			}()
			go func() {
				defer func() {
					if r := recover(); r != nil {
						panic_fmt.RecoverError(os.Getenv("RAW_LOG") != "true", r)
					}
				}()

				_, msg, err := conn.ReadMessage() // Read the next message from the client.
				if err != nil {
					return
				}
				hrpcMsg, err := hrpc.UnmarshalHRPC(msg, "application/hrpc", h.Input)
				if err != nil {
					conn.Close() // TODO: don't close socket, send error to client
					return
				}
				reqChan <- hrpcMsg
			}()

			for msg := range sendChan {
				resp, err := hrpc.MarshalHRPC(msg, "application/hrpc")
				if err != nil {
					conn.Close()
					return
				}
				conn.WriteMessage(websocket.BinaryMessage, append([]byte{0}, resp...))
			}
		}, websocket.Config{
			Subprotocols: []string{"hrpc1"},
		})(c)
	}
}
