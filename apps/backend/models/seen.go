package models

import (
	"time"

	"gorm.io/gorm"
)

type Seen struct {
	gorm.Model
	Progress   float32
	FinshedAt  time.Time
	UserID     uint
	User       User
	MetadataID uint
	Metadata   Metadata
}
