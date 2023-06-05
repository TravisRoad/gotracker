package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Rating  float32
	Content string
	Spoiler bool
	// Visibility string // TODO: 或许指定用户可见?
	UserID     uint
	User       User
	MetadataID uint
	Metadata   Metadata
}
