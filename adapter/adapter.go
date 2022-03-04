// Package adapter is an adapter for the hrpc protocol.
package adapter

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	goserver "github.com/harmony-development/hrpc/pkg/go-server"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func RegisterHandler[T context.Context](router fiber.Router, handler goserver.MethodHandler[T]) {
	for route, h := range handler.Routes() {
		router.All(route, UnaryHandler(h))
	}
	for route, h := range handler.StreamRoutes() {
		router.All(route, StreamHandler(h))
	}
}

func UnaryHandler[T context.Context](h goserver.UnaryMethodData[T]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		input, err := UnmarshalHRPC(c.Body(), c.Get("Content-Type"), h.Input)
		if err != nil {
			return err
		}
		result, err := h.Handler(c.UserContext().(T), input)
		if err != nil {
			return err
		}
		response, err := MarshalHRPC(result, c.Get("Content-Type"))
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
				if err := h.Handler(ctx.(T), reqChan, sendChan); err != nil {
					close(reqChan)
					close(sendChan)
				}
			}()
			go func() {
				_, msg, err := conn.ReadMessage() // Read the next message from the client.
				if err != nil {
					conn.Close()
					return
				}
				hrpcMsg, err := UnmarshalHRPC(msg, "application/hrpc", h.Input)
				if err != nil {
					conn.Close() // TODO: don't close socket, send error to client
					return
				}
				reqChan <- hrpcMsg
			}()

			for msg := range sendChan {
				resp, err := MarshalHRPC(msg, "application/hrpc")
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

// MarshalHRPC marshals a message to the hrpc protocol.
func MarshalHRPC(content goserver.VTProtoMessage, contentType string) ([]byte, error) {
	var response []byte

	var err error

	switch contentType {
	case "application/hrpc-json":
		response, err = protojson.Marshal(content)
	default:
		response, err = content.MarshalVT()
	}

	return response, err
}

func UnmarshalHRPC[T goserver.VTProtoMessage](content []byte, contentType string, messageType T) (T, error) {
	newMessage := proto.Clone(messageType).(T)

	var err error

	switch contentType {
	case "application/hrpc-json":
		if err := protojson.Unmarshal(content, newMessage); err != nil {
			return *new(T), err
		}
	default:
		if err := proto.Unmarshal(content, newMessage); err != nil {
			return *new(T), err
		}
	}
	return newMessage, err
}
