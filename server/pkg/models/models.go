package models

import "database/sql"

type Models struct {
	Feeds     FeedModel
	FeedItems FeedItemModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Feeds:     FeedModel{DB: db},
		FeedItems: FeedItemModel{DB: db},
	}
}
