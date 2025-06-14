package auth

import (
	"net/mail"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/jmoiron/sqlx"

	"creditninja/internal/models"
	"creditninja/internal/services"
)

func Register(c *fiber.Ctx) error {
	type form struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	var data form
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString("Bad input")
	}
	if _, err := mail.ParseAddress(data.Email); err != nil {
		return c.Status(400).SendString("Invalid email")
	}
	db := c.Locals("db").(*sqlx.DB)
	user, err := models.CreateUser(db, data.Email, data.Password, "client")
	if err != nil {
		c.Context().Logger().Printf("User creation failed: %v", err)
		return c.Status(400).Render("login", fiber.Map{"Register": true, "Error": "Failed to create user. This email might already be registered."})
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

func Login(c *fiber.Ctx) error {
	type form struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	var data form
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString("Bad input")
	}
	db := c.Locals("db").(*sqlx.DB)
	user, err := models.GetUserByEmail(db, data.Email)
	if err != nil || !user.CheckPassword(data.Password) {
		return c.Status(401).Render("login", fiber.Map{"Error": "Invalid credentials"})
	}
	if !user.Verified {
		return c.Status(401).Render("login", fiber.Map{"Error": "Please check your email and verify your account first."})
	}
	store := c.Locals("store").(*session.Store)
	sess, _ := store.Get(c)
	sess.Set("user_id", user.ID.String())
	sess.Save()
	return c.Redirect("/dashboard")
}

func Logout(c *fiber.Ctx) error {
	store := c.Locals("store").(*session.Store)
	sess, _ := store.Get(c)
	sess.Destroy()
	return c.Redirect("/login")
}
