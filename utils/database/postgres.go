package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

var DB *pgxpool.Pool

func ConnectToDatabase() error {
	db, err := pgxpool.Connect(context.Background(), os.Getenv("DB_URI"))
	if err != nil {
		return err
	}

	DB = db
	log.Println("Connected to PostgreSQL.")

	return nil
}
