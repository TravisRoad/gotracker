package models

import "gorm.io/gorm"

type Manga struct {
	gorm.Model
	MetadataID uint
	Metadata   Metadata
}
