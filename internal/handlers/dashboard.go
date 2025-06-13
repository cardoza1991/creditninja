package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"creditninja/internal/models"
)

func dashboard(c *fiber.Ctx) error {
	store := c.Locals("store").(*session.Store)
	sess, _ := store.Get(c)
	uidRaw := sess.Get("user_id")
	if uidRaw == nil {
		return c.Redirect("/login")
	}
	userID, _ := uuid.Parse(uidRaw.(string))
	db := c.Locals("db").(*sqlx.DB)

	var reports []models.CreditReport
	db.Select(&reports, "SELECT * FROM credit_reports WHERE user_id=$1 ORDER BY created_at DESC", userID)

	return c.Render("dashboard", fiber.Map{
		"Reports": reports,
	})
}
