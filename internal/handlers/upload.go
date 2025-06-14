package handlers

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"creditninja/internal/models"
	"creditninja/internal/services"
)

func upload(c *fiber.Ctx) error {
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
	if !user.Verified {
		return c.Redirect("/check-email?email=" + user.Email)
	}
	if !user.Paid {
		return c.Status(fiber.StatusPaymentRequired).SendString("Upgrade required")
	}

	fileHeader, err := c.FormFile("report")
	if err != nil {
		return c.Status(400).SendString("File required")
	}

	saveDir := "./static/uploads"
	os.MkdirAll(saveDir, 0755)
	fname := uuid.New().String() + filepath.Ext(fileHeader.Filename)
	rawPath := filepath.Join(saveDir, fname)
	if err := c.SaveFile(fileHeader, rawPath); err != nil {
		return c.Status(500).SendString("Upload failed")
	}

	parsedReport, err := services.ParseReport(rawPath)
	if err != nil {
		c.Context().Logger().Printf("parse error: %v", err)
	}
	parsed, _ := json.Marshal(parsedReport)

	_, err = models.CreateReport(db, userID, rawPath, string(parsed))
	if err != nil {
		return c.Status(500).SendString("DB error")
	}
	return c.Redirect("/dashboard")
}
