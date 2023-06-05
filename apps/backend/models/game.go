package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	MetadataID uint
	Metadata   Metadata
}
