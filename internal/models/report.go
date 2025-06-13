package models

import (
    "time"

    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
)

type CreditReport struct {
    ID         uuid.UUID `db:"id"`
    UserID     uuid.UUID `db:"user_id"`
    RawPath    string    `db:"raw_path"`
    ParsedJSON string    `db:"parsed_json"`
    CreatedAt  time.Time `db:"created_at"`
}

func CreateReport(db *sqlx.DB, userID uuid.UUID, path, parsed string) (*CreditReport, error) {
    rep := &CreditReport{
        ID: uuid.New(),
        UserID: userID,
        RawPath: path,
        ParsedJSON: parsed,
        CreatedAt: time.Now(),
    }
    _, err := db.NamedExec(`INSERT INTO credit_reports (id, user_id, raw_path, parsed_json, created_at)
                             VALUES (:id, :user_id, :raw_path, :parsed_json, :created_at)`, rep)
    return rep, err
}
