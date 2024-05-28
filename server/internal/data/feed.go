package data

import (
	"context"
	"database/sql"
	"time"
)

type Feed struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
}

type FeedModel struct {
	DB *sql.DB
}

func (m FeedModel) Insert(feed *Feed) error {
	query := `INSERT INTO feeds (title, url) VALUES ($1, $2) RETURNING id`

	args := []any{feed.Title, feed.URL}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&feed.ID)
}
