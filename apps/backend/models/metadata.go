package models

import (
	"gorm.io/gorm"
)

type Metadata struct {
	gorm.Model
	Title         string
	Description   string
	PublishYear   string
	PublishData   string
	ImageUrl      string
	Creators      string
	SourceUrl     string
	TitleCN       string
	DescriptionCN string
	ImageUrlCN    string
	CreatorsCN    string
	Keywords      []*Keyword `gorm:"many2many:metadata_keywords;"`
	Collections   []*Keyword `gorm:"many2many:collection_metadata;"`
}
