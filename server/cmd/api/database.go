package main

import (
	"ajbates93/rss-feed/pkg/models"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// Initialise the database connection
	dsn := os.Getenv("RSS_FEED_DSN")
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}

	// For some reason this is not working...
	db.AutoMigrate(&models.FeedModel{}, &models.FeedItemModel{})

	return db, nil
}
