package models

import "gorm.io/gorm"

type Keyword struct {
	gorm.Model
	Name      string
	Metadatas []*Metadata `gorm:"many2many:metadata_keywords;"`
}
