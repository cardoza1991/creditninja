package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"creditninja/internal/models"
	"creditninja/internal/services"
)

func VerifyEmail(c *fiber.Ctx) error {
	token := c.Params("token")
	db := c.Locals("db").(*sqlx.DB)

	vt, err := models.GetVerificationToken(db, token)
	if err != nil {
		return c.Status(400).SendString("Invalid token")
	}

	_, err = db.Exec(`UPDATE users SET verified=true WHERE id=$1`, vt.UserID)
	if err != nil {
		return c.Status(500).SendString("DB error")
	}

	models.DeleteVerificationToken(db, token)
	return c.Redirect("/verified")
}

// ResendVerification generates a new token and emails it to the user.
func ResendVerification(c *fiber.Ctx) error {
	email := c.FormValue("email")
	if email == "" {
		return c.Status(400).SendString("Email required")
	}
	db := c.Locals("db").(*sqlx.DB)
	user, err := models.GetUserByEmail(db, email)
	if err != nil {
		return c.Status(400).SendString("User not found")
	}
	vt, err := models.CreateVerificationToken(db, user.ID)
	if err == nil {
		appURL := os.Getenv("APP_URL")
		if appURL == "" {
			appURL = "http://localhost:8080"
		}
		link := appURL + "/verify/" + vt.Token
		services.SendEmail(user.Email, "Verify your email", link)
	}
	return c.Redirect("/check-email?email=" + user.Email)
}
