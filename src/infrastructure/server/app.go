package infraestructure

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	config "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/config"
	types "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/types"
	middleware "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/http"
	"go.uber.org/fx"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setRoutesByModule(app *fiber.App, h *types.HandlersStore) {
	for _, handlerModule := range h.Handlers {
		route := app.Group("/" + handlerModule.Prefix)
		app.Static("/downloads", "./uploads")
		for _, routeItem := range handlerModule.Routes {
			// Configurar la ruta específica
			switch routeItem.Method {
			case http.MethodGet:
				if routeItem.RequiresAuth {
					route.Get(routeItem.Route, middleware.JWTMiddleware, routeItem.Handler)
				} else {
					route.Get(routeItem.Route, routeItem.Handler)
				}
			case http.MethodPost:
				if routeItem.RequiresAuth {
					route.Post(routeItem.Route, middleware.JWTMiddleware, routeItem.Handler)
				} else {
					route.Post(routeItem.Route, routeItem.Handler)
				}
			case http.MethodPut:
				if routeItem.RequiresAuth {
					route.Put(routeItem.Route, middleware.JWTMiddleware, routeItem.Handler)
				} else {
					route.Put(routeItem.Route, routeItem.Handler)
				}
			case http.MethodDelete:
				if routeItem.RequiresAuth {
					route.Delete(routeItem.Route, middleware.JWTMiddleware, routeItem.Handler)
				} else {
					route.Delete(routeItem.Route, routeItem.Handler)
				}
			case http.MethodPatch:
				if routeItem.RequiresAuth {
					route.Patch(routeItem.Route, middleware.JWTMiddleware, routeItem.Handler)
				} else {
					route.Patch(routeItem.Route, routeItem.Handler)
				}
			}
		}
	}
}
func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"isError": true,
		"message": err.Error(),
	})
}
func NewHttpFiberServer(lc fx.Lifecycle, h *types.HandlersStore, cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	setRoutesByModule(app, h)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting fiber server on port " + cfg.App.Port)
			go app.Listen(":" + cfg.App.Port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}
