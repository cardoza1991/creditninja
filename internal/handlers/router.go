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

	app.Get("/check-email", func(c *fiber.Ctx) error {
		return c.Render("check_email", fiber.Map{"Email": c.Query("email")})
	})
	app.Get("/verified", func(c *fiber.Ctx) error { return c.Render("verified", nil) })
	app.Post("/resend-verification", ResendVerification)
	app.Get("/verify/:token", VerifyEmail)

	app.Get("/dashboard", dashboard)

	app.Post("/upload", upload)

	app.Get("/pay", pay)
}
