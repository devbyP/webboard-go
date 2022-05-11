package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:32"`
	Hash     string `gorm:"size:"`
	Name     string `gorm:"size:128"`
	Email    string
	Bio      string `gorm:"size:2048"`
	Posts    []Post
	Reps     []Rep
}
