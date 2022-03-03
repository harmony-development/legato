package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/harmony-development/legato/adapter"
	"github.com/harmony-development/legato/api"
	authv1impl "github.com/harmony-development/legato/api/auth/v1"
	"github.com/harmony-development/legato/config"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
)

// Server is an instance of the server.
type Server struct {
	app    *fiber.App
	config *config.Config
}

// New creates a new server instance.
func New(config *config.Config) *Server {
	app := fiber.New()

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	if config.UseLocalCORS {
		app.Use(cors.New())
	}

	app.Use(func(c *fiber.Ctx) error {
		c.SetUserContext(api.LegatoContext{})
		return c.Next()
	})

	authService := &authv1.AuthServiceHandler[api.LegatoContext]{
		Impl: &authv1impl.AuthV1{},
	}

	adapter.RegisterHandler[api.LegatoContext](app, authService)

	return &Server{
		app,
		config,
	}
}

// Start starts the server.
func (s *Server) Start() error {
	return s.app.Listen(fmt.Sprintf(":%d", s.config.Port))
}
