package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"creditninja/internal/models"
)

func billing(c *fiber.Ctx) error {
	store := c.Locals("store").(*session.Store)
	sess, _ := store.Get(c)
	uidRaw := sess.Get("user_id")
	if uidRaw == nil {
		return c.Redirect("/login")
	}
	userID, _ := uuid.Parse(uidRaw.(string))

	db := c.Locals("db").(*sqlx.DB)
	user, err := models.GetUserByID(db, userID)
	if err != nil {
		return c.Redirect("/login")
	}

	return c.Render("billing", fiber.Map{"User": user})
}
