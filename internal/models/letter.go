package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type DisputeLetter struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	ReportID  uuid.UUID `db:"report_id"`
	PdfPath   string    `db:"pdf_path"`
	Round     int       `db:"round"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}

func CreateLetter(db *sqlx.DB, letter *DisputeLetter) error {
	_, err := db.NamedExec(`INSERT INTO dispute_letters (id, user_id, report_id, pdf_path, round, status, created_at)
                             VALUES (:id, :user_id, :report_id, :pdf_path, :round, :status, :created_at)`, letter)
	return err
}
