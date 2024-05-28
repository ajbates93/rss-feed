package rss

import (
	"ajbates93/rss-feed/pkg/models"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
	"gorm.io/gorm"
)

func FetchFeed(ctx context.Context, feedURLFromClient string) (*gofeed.Feed, error) {
	log.Printf("Fetching feed %s", feedURLFromClient)
	if feedURLFromClient == "" {
		return nil, fmt.Errorf("empty feed URL")
	}

	parser := gofeed.NewParser()
	feed, err := parser.ParseURLWithContext(feedURLFromClient, ctx)
	if err != nil {
		return nil, err
	}
	return feed, nil
}

func SaveFeed(ctx context.Context, db *gorm.DB, feed *gofeed.Feed) error {
	log.Printf("Saving feed %s", feed.Title)

	// Check if the feed already exists
	var existingFeed models.FeedModel

	if err := db.Where("url = ?", feed.FeedLink).First(&existingFeed).Error; err != nil {
		// Feed already exists, skip insertion
		if err != gorm.ErrRecordNotFound {
			return err
		}

		// Feed does not exist, create a new one
		// Save the feed to the database
		dbFeed := models.FeedModel{
			Title: feed.Title,
			URL:   feed.FeedLink,
		}

		if err := db.Create(&dbFeed).Error; err != nil {
			return err
		}

		existingFeed = dbFeed
	}

	log.Printf("Feed %s exists, updating feed items", feed.Title)

	// Save the feed items to the database
	for _, item := range feed.Items {
		dbItem := &models.FeedItemModel{
			FeedID:      existingFeed.ID,
			Author:      existingFeed.Title,
			Title:       item.Title,
			Description: item.Description,
			URL:         item.Link,
		}

		if item.PublishedParsed != nil {
			dbItem.PublishedAt = *item.PublishedParsed
		}
		if item.Image != nil {
			dbItem.ImageURL = item.Image.URL
		}

		if err := db.WithContext(ctx).Create(dbItem).Error; err != nil {
			if err.Error() == "ERROR: duplicate key value violates unique constraint \"idx_feed_item_url\"" {
				log.Printf("Feed item %s already exists, skipping", dbItem.Title)
				continue
			}
			return err
		}
		log.Printf("Feed item %s saved", dbItem.Title)
	}

	// Check the total count of articles in the database
	var count int64
	if err := db.Model(&models.FeedItemModel{}).Count(&count).Error; err != nil {
		return err
	}

	// If the count exceeds the threshold, remove the oldest articles
	if count > 100 {
		if err := db.Where("id IN (?)", db.Model(&models.FeedItemModel{}).Order("published_at ASC").Limit(int(count-100)).Select("id")).Delete(&models.FeedItemModel{}).Error; err != nil {
			return err
		}
	}

	return nil
}

func StartFeedFetcher(db *gorm.DB, interval time.Duration) {
	log.Printf("Starting feed fetcher with interval %v", interval)

	// Initial execution of fetch and save
	FetchAndSaveFeeds(db)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		FetchAndSaveFeeds(db)
	}
}

func FetchAndSaveFeeds(db *gorm.DB) error {
	feeds := []models.FeedModel{}
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
			return err
		}
	}

	if len(feeds) == 0 {
		log.Println("No feeds to fetch")
	}

	return nil
}
