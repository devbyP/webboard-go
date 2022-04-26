package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title     string `gorm:"size:256"`
	Content   string `gorm:"size:1024"`
	Upvotes   int
	Downvotes int
	UserID    uint
	Reps      []Rep
}
