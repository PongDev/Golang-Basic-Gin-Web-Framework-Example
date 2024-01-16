package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title   string `gorm:"unique"`
	Content string
}
