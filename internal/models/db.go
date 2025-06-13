package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(dsn string) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", dsn)
}
