package models

import "gorm.io/gorm"

type Feed struct {
	gorm.Model
	Title string
	URL   string
}
