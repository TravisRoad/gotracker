package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	TextNum    int
	PageNum    int
	MetadataID uint
	Metadata   Metadata
}
