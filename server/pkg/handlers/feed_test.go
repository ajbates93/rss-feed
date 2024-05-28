package handlers_test

import (
	"ajbates93/rss-feed/pkg/handlers"
	"ajbates93/rss-feed/pkg/models"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateFeed(t *testing.T) {
	// Set up the test database connection
	db := setupTestDB()

	// Create a new instance of RSSFeedHandler
	feedHandler := &handlers.RSSFeedHandler{DB: db}

	// Create a new feed payload
	feed := &models.FeedModel{
		Title: "Test Feed",
		URL:   "http://example.com/feed",
	}
	payload, err := json.Marshal(feed)
	require.NoError(t, err)

	// Create a new request
	req, err := http.NewRequest("POST", "/feeds", bytes.NewBuffer(payload))
	require.NoError(t, err)

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the CreateFeed handler
	handler := http.HandlerFunc(feedHandler.CreateFeed)
	handler.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var createdFeed models.FeedModel
	err = json.NewDecoder(rr.Body).Decode(&createdFeed)
	require.NoError(t, err)
	assert.Equal(t, feed.Title, createdFeed.Title)
	assert.Equal(t, feed.URL, createdFeed.URL)
	assert.NotEmpty(t, createdFeed.ID)
}

func setupTestDB() *gorm.DB {
	// Set up the test database connection
	db, err := gorm.Open(postgres.Open("postgres://alexbates:!Luc0zad32648@localhost/rss_feed?sslmode=disable"))
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.FeedModel{}, &models.FeedItemModel{})
	return db
}
