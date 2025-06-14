package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type VerificationToken struct {
	Token     string    `db:"token"`
	UserID    uuid.UUID `db:"user_id"`
	ExpiresAt time.Time `db:"expires_at"`
}

func CreateVerificationToken(db *sqlx.DB, userID uuid.UUID) (*VerificationToken, error) {
	vt := &VerificationToken{
		Token:     uuid.New().String(),
		UserID:    userID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	_, err := db.NamedExec(`INSERT INTO verification_tokens (token, user_id, expires_at)
                            VALUES (:token, :user_id, :expires_at)`, vt)
	return vt, err
}

func GetVerificationToken(db *sqlx.DB, token string) (*VerificationToken, error) {
	var vt VerificationToken
	err := db.Get(&vt, `SELECT * FROM verification_tokens WHERE token=$1`, token)
	return &vt, err
}

func DeleteVerificationToken(db *sqlx.DB, token string) error {
	_, err := db.Exec(`DELETE FROM verification_tokens WHERE token=$1`, token)
	return err
}
