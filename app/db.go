package app

import (
	"TEST2/app/core"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"os"
)

var (
	DB *sql.DB
)

func connectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
}

// InitDatabase For Vanilla SQL
func InitDatabase(log core.Logger) (*sql.DB, error) {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), connectionString())
	if err != nil {
		return nil, err
	}
	log.Info("Database connection was created")
	DB = db
	return db, nil
}

func RunMigrations() error {
	err := goose.Up(DB, "./migrations")
	if err != nil {
		return err
	}
	return nil
}
