package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "auth-practice/docs" // docs is generated by Swag CLI, you have to import it.
)

func AddSwaggerRoutes(app *fiber.App) {
	// setup swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
}
