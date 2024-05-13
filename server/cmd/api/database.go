package main

import (
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Feed struct {
	gorm.Model
	Title string
	URL   string
}

func InitDB() (*gorm.DB, error) {
	// Initialise the database connection
	dsn := os.Getenv("RSS_FEED_DSN")
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Feed{})

	return db, nil
}
