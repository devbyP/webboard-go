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

func (p *Post) GetAllPosts(db *gorm.DB, limit, offset int) ([]*Post, error) {
	rows, err := db.Model(p).Order("created_at").Limit(limit).Offset(offset).Rows()
	if err != nil {
		return nil, err
	}
	posts := make([]*Post, 0)
	for rows.Next() {
		post := Post{}
		if err = db.ScanRows(rows, &post); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}
