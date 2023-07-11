package models

import (
	"time"

	"gorm.io/gorm"
)

type Seen struct {
	gorm.Model
	Status     uint
	Progress   uint
	FinshedAt  time.Time
	Identifier string
	Source     string
	Variety    string
	UserID     uint
	User       User
	MetadataID uint
	Metadata   Metadata
}

func (s *Seen) Save() error {
	if err := DB.Save(s).Error; err != nil {
		return err
	}
	return nil
}
