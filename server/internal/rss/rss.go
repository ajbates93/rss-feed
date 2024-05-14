package rss

import (
	"ajbates93/rss-feed/pkg/models"
	"context"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
	"gorm.io/gorm"
)

func FetchFeed(ctx context.Context, feedURL string) (*gofeed.Feed, error) {
	log.Printf("Fetching feed %s", feedURL)
	parser := gofeed.NewParser()
	feed, err := parser.ParseURLWithContext(feedURL, ctx)
	if err != nil {
		return nil, err
	}
	return feed, nil
}

func SaveFeed(ctx context.Context, db *gorm.DB, feed *gofeed.Feed) error {
	log.Printf("Saving feed %s", feed.Title)
	// Save the feed to the database
	dbFeed := models.Feed{
		Title: feed.Title,
		URL:   feed.FeedLink,
	}

	if err := db.Create(&dbFeed).Error; err != nil {
		return err
	}

	// Save the feed items to the database
	for _, item := range feed.Items {
		dbItem := &models.FeedItem{
			FeedID:      dbFeed.ID,
			Author:      dbFeed.Title,
			Title:       item.Title,
			Description: item.Description,
			URL:         item.Link,
			PublishedAt: *item.PublishedParsed,
		}
		if err := db.WithContext(ctx).Create(dbItem).Error; err != nil {
			return err
		}
	}

	return nil
}

func StartFeedFetcher(db *gorm.DB, interval time.Duration) {
	log.Printf("Starting feed fetcher with interval %v", interval)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		feeds := []models.Feed{}
		db.Find(&feeds)

		for _, feed := range feeds {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			parsedFeed, err := FetchFeed(ctx, feed.URL)
			if err != nil {
				log.Printf("Error fetching feed %s: %v", feed.URL, err)
				continue
			}

			if err := SaveFeed(ctx, db, parsedFeed); err != nil {
				log.Printf("Error saving feed %s: %v", feed.URL, err)
			}
		}
	}
}
