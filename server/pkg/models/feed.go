package models

import (
	"time"

	"gorm.io/gorm"
)

type Feed struct {
	gorm.Model
	Title string `gorm:"not null"`
	URL   string `gorm:"not null;uniqueIndex"`
}

type FeedItem struct {
	gorm.Model
	FeedID      uint   `gorm:"not null"`
	Author      string `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string
	URL         string `gorm:"not null"`
	PublishedAt time.Time
}
