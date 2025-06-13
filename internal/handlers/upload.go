package handlers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"creditninja/internal/models"
)

func upload(c *fiber.Ctx) error {
	store := c.Locals("store").(*session.Store)
	sess, _ := store.Get(c)
	uidRaw := sess.Get("user_id")
	if uidRaw == nil {
		return c.Redirect("/login")
	}
	userID, _ := uuid.Parse(uidRaw.(string))

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

	// Parse stub
	parsed, _ := json.Marshal(map[string]string{
		"status":   "placeholder",
		"uploaded": time.Now().String(),
	})

	db := c.Locals("db").(*sqlx.DB)
	_, err = models.CreateReport(db, userID, rawPath, string(parsed))
	if err != nil {
		return c.Status(500).SendString("DB error")
	}
	return c.Redirect("/dashboard")
}
