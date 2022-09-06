package server

import (
	"fmt"
	"strings"

	"github.com/shreeyashnaik/Course-Management-Backend/src/core/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/shreeyashnaik/Course-Management-Backend/config"
)

func healthCheck(ctx *fiber.Ctx) error {
	return ctx.SendString("OK")
}

func StartCoreServer() {

	// Create Fiber App
	app := fiber.New(fiber.Config{
		AppName:       "Course Management Server",
		StrictRouting: true,
	})

	// Health Check routers
	app.Get("/ok", healthCheck)
	app.Post("/", healthCheck)

	app.Use(logger.New(logger.Config{Next: func(c *fiber.Ctx) bool {
		// Log all routers except healthcheck
		return strings.HasPrefix(c.Path(), "api")
	}}))

	// Initialise CORS
	app.Use(cors.New(cors.Config{
		// Allowed Origins are based on environment.
		AllowOrigins:     config.CORS_ALLOWED_ORIGINS,
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))

	// Mount routes for the above app
	routers.MountRoutes(app)

	// Start Fiber Server
	err := app.Listen(fmt.Sprintf(":%s", config.PORT))
	if err != nil {
		panic(err)
	}
}
