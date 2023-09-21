package ports

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type WebServer interface {
	Start() error
	Shutdown(ctx context.Context) error
	Use(args ...any) fiber.Router
	Group(prefix string, handlers ...fiber.Handler) fiber.Router
}
