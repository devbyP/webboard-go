package models

import (
	"gorm.io/gorm"
)

type Rep struct {
	gorm.Model
	Content   string `gorm:"size:1024"`
	Upvotes   int
	Downvotes int
	UserID    uint
	PostID    uint
}
