package handlers

import (
	"ajbates93/rss-feed/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type RSSFeedItemHandler struct {
	DB *gorm.DB
}

func (h *RSSFeedItemHandler) GetFeedItems(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	var feedItems []models.FeedItem
	result := h.DB.Order("published_at desc").Limit(limit).Offset(offset).Find(&feedItems)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	type ResponseItem struct {
		Title       string `json:"title"`
		Author      string `json:"author"`
		URL         string `json:"url"`
		PublishedAt string `json:"publishedAt"`
		ImageURL    string `json:"imageURL"`
	}

	var responseItems []ResponseItem
	for _, item := range feedItems {
		formattedPublishedAt := item.PublishedAt.Format("02/01/2006")
		responseItems = append(responseItems, ResponseItem{
			Title:       item.Title,
			Author:      item.Author,
			PublishedAt: formattedPublishedAt,
			URL:         item.URL,
			ImageURL:    item.ImageURL,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":    responseItems,
		"success": true,
		"meta": map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": len(feedItems),
		},
	})
}

func (h *RSSFeedItemHandler) RegisterRoutes(r chi.Router) {
	r.Get("/feed-items", h.GetFeedItems)
}
