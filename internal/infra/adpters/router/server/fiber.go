package server

import (
	"context"
	"fmt"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"strconv"
)

const (
	empty            = ""
	slash            = "/"
	defaultPort      = "8080"
	defaultBaseGroup = slash + "api"
	errorStartingApp = "error starting app: %w"
	errorStoppingApp = "error stopping app: %w"
)

type Server struct {
	app    *fiber.App
	router fiber.Router
	config *types.SvrConfig
}

func NewFiberServer(serverConfig ...types.SvrConfig) *Server {
	var config types.SvrConfig
	var routesGroup fiber.Router
	var pathBase bool
	var pathName string

	// Start App
	app := fiber.New(fiber.Config{
		ErrorHandler: CustomErrorHandler,
	})

	// Initialize default config (Assign the middleware to /metrics)
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics Services"}))

	// ---------------------------------------------------------------------------------------------------------- //

	if len(serverConfig) >= 1 {
		config = serverConfig[0]
	}

	if pathBase, pathName = getBaseGroup(config); pathBase || (!pathBase && pathName != empty) {
		routesGroup = app.Group(fmt.Sprintf("/%s", pathName)).Name(fmt.Sprintf("%s-group", pathName))
	} else {
		routesGroup = app.Group(slash).Name("main-group")
	}

	// ---------------------------------------------------------------------------------------------------------- //
	return &Server{
		app:    app,
		router: routesGroup,
		config: &config,
	}
}

func (fb *Server) Start() error {
	port := defaultPort
	if fb.config != nil {
		if fb.config.Port > 0 {
			port = strconv.Itoa(int(fb.config.Port))
		}
	}

	if err := fb.app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		return fmt.Errorf(errorStartingApp, err)
	}
	return nil
}

func (fb *Server) Shutdown(_ context.Context) error {
	if err := fb.app.Shutdown(); err != nil {
		return fmt.Errorf(errorStoppingApp, err)
	}
	return nil
}

func (fb *Server) Use(args ...any) fiber.Router {
	return fb.app.Use(args...)
}

func (fb *Server) Group(prefix string, handlers ...fiber.Handler) fiber.Router {
	return fb.router.Group(prefix, handlers...)
}

// --------------------------------------------------------------------------------------------------------------- //

func getBaseGroup(config types.SvrConfig) (pb bool, pn string) {
	if config.PathBase != nil {
		pb = *config.PathBase
	}

	if config.PathName != nil {
		if *config.PathName != empty && *config.PathName != slash {
			pn = *config.PathName
		}
	}

	if pb {
		if pb && pn == empty {
			pn = defaultBaseGroup
		}
	}
	return
}
