package models

import (
	"database/sql"
	"time"
)

type FeedItem struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
	FeedID      uint      `json:"feedID"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	PublishedAt time.Time `json:"publishedAt"`
	ImageURL    string    `json:"imageURL"`
}

type FeedItemModel struct {
	DB *sql.DB
}
