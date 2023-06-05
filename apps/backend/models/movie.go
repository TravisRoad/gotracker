package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Time       int
	MetadataID uint
	Metadata   Metadata
}
