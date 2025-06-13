package handlers

import (
	"github.com/gofiber/fiber/v2"

	"creditninja/internal/auth"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error { return c.Render("login", nil) })
	app.Get("/register", func(c *fiber.Ctx) error { return c.Render("login", fiber.Map{"Register": true}) })
	app.Post("/register", auth.Register)
	app.Get("/login", func(c *fiber.Ctx) error { return c.Render("login", nil) })
	app.Post("/login", auth.Login)
	app.Get("/logout", auth.Logout)

	app.Get("/dashboard", dashboard)

	app.Get("/pay", pay)
}
