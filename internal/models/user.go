package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	Verified  bool      `db:"verified"`
	Paid      bool      `db:"paid"`
	CreatedAt time.Time `db:"created_at"`
}

func (u *User) SetPassword(raw string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func (u *User) CheckPassword(raw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(raw)) == nil
}

func CreateUser(db *sqlx.DB, email, password, role string) (*User, error) {
	user := &User{
		ID:        uuid.New(),
		Email:     email,
		Role:      role,
		Verified:  false,
		Paid:      false,
		CreatedAt: time.Now(),
	}
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}
	_, err := db.NamedExec(`INSERT INTO users (id, email, password, role, verified, paid, created_at)
                             VALUES (:id, :email, :password, :role, :verified, :paid, :created_at)`, user)
	return user, err
}

func GetUserByEmail(db *sqlx.DB, email string) (*User, error) {
	var u User
	err := db.Get(&u, "SELECT * FROM users WHERE email=$1", email)
	return &u, err
}

func GetUserByID(db *sqlx.DB, id uuid.UUID) (*User, error) {
	var u User
	err := db.Get(&u, "SELECT * FROM users WHERE id=$1", id)
	return &u, err
}
