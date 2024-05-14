package handlers

import (
	"ajbates93/rss-feed/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type RSSFeedHandler struct {
	DB *gorm.DB
}

func (h *RSSFeedHandler) CreateFeed(w http.ResponseWriter, r *http.Request) {
	var feed models.Feed
	err := json.NewDecoder(r.Body).Decode(&feed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.DB.Create(&feed).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}

func (h *RSSFeedHandler) GetAllFeeds(w http.ResponseWriter, r *http.Request) {
	var feeds []models.Feed
	if err := h.DB.Find(&feeds).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feeds)
}

func (h *RSSFeedHandler) UpdateFeed(w http.ResponseWriter, r *http.Request) {
	var feed models.Feed
	err := json.NewDecoder(r.Body).Decode(&feed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.DB.Save(&feed).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}

func (h *RSSFeedHandler) DeleteFeed(w http.ResponseWriter, r *http.Request) {
	// Get ID from url params
	id := chi.URLParam(r, "id")

	// Find feed by ID
	var feed models.Feed
	if err := h.DB.First(&feed, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Delete feed
	if err := h.DB.Delete(&feed).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}

func (h *RSSFeedHandler) RegisterRoutes(r chi.Router) {
	r.Get("/feeds", h.GetAllFeeds)
	r.Post("/feeds", h.CreateFeed)
	r.Put("/feeds", h.UpdateFeed)
	r.Delete("/feeds/{id}", h.DeleteFeed)
}
