package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/harmony-development/legato/adapter"
	"github.com/harmony-development/legato/api"
	authv1impl "github.com/harmony-development/legato/api/auth/v1"
	chatv1impl "github.com/harmony-development/legato/api/chat/v1"
	"github.com/harmony-development/legato/api/middleware"
	profilev1impl "github.com/harmony-development/legato/api/profile/v1"
	"github.com/harmony-development/legato/config"
	"github.com/harmony-development/legato/db"
	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	profilev1 "github.com/harmony-development/legato/gen/profile/v1"
)

// Server is an instance of the server.
type Server struct {
	app    *fiber.App
	config *config.Config
}

// New creates a new server instance.
func New(config *config.Config) (*Server, error) {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(Recover(config.CleanLogs))

	if config.UseLocalCORS {
		app.Use(cors.New(cors.Config{
			MaxAge: 9600,
		}))
	}

	app.Use(func(c *fiber.Ctx) error {
		c.SetUserContext(&api.LegatoContext{
			Context: c.Context(),
		})
		return c.Next()
	})

	db, err := db.New(context.TODO(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	middlewareHandler := middleware.NewHandler(db)
	app.Use(middlewareHandler.ErrorHandler())

	authService := &authv1.AuthServiceHandler[*api.LegatoContext]{
		Impl: authv1impl.New(db, config),
	}

	chatService := &chatv1.ChatServiceHandler[*api.LegatoContext]{
		Impl: chatv1impl.New(db),
	}

	profileService := &profilev1.ProfileServiceHandler[*api.LegatoContext]{
		Impl: profilev1impl.New(db),
	}

	adapter.RegisterHandler[*api.LegatoContext](middlewareHandler, app, authService)
	adapter.RegisterHandler[*api.LegatoContext](middlewareHandler, app, chatService)
	adapter.RegisterHandler[*api.LegatoContext](middlewareHandler, app, profileService)

	return &Server{
		app,
		config,
	}, nil
}

// Start starts the server.
func (s *Server) Start() error {
	return s.app.Listen(fmt.Sprintf(":%d", s.config.Port))
}
