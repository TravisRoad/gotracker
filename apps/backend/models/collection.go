package models

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Name      string
	UserID    uint
	User      User
	Metadatas []*Metadata `gorm:"many2many:collection_metadata;"`
}
