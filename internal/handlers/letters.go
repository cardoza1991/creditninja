package handlers

import (
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"creditninja/internal/models"
	"creditninja/internal/services"
)

func createLetter(c *fiber.Ctx) error {
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

	reportIDStr := c.FormValue("report_id")
	if reportIDStr == "" {
		return c.Status(400).SendString("report_id required")
	}
	reportID, err := uuid.Parse(reportIDStr)
	if err != nil {
		return c.Status(400).SendString("bad report_id")
	}

	var report models.CreditReport
	if err := db.Get(&report, "SELECT * FROM credit_reports WHERE id=$1 AND user_id=$2", reportID, userID); err != nil {
		return c.Status(404).SendString("report not found")
	}

	prompt := "Generate a dispute letter based on this report: " + report.ParsedJSON
	text, err := services.GenerateLetter(prompt)
	if err != nil {
		return c.Status(500).SendString("AI error")
	}

	outDir := "./static/letters"
	os.MkdirAll(outDir, 0755)
	pdfPath := filepath.Join(outDir, uuid.New().String()+".pdf")
	if err := services.RenderPDF(text, pdfPath); err != nil {
		return c.Status(500).SendString("PDF error")
	}

	letter := &models.DisputeLetter{
		ID:        uuid.New(),
		UserID:    userID,
		ReportID:  reportID,
		PdfPath:   pdfPath,
		Round:     1,
		Status:    "pending",
		CreatedAt: time.Now(),
	}
	if err := models.CreateLetter(db, letter); err != nil {
		return c.Status(500).SendString("DB error")
	}

	return c.Redirect("/dashboard")
}
