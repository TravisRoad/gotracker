package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Rating      float32
	ExtraRating []byte `gorm:"type:BLOB"`
	Content     string
	// Spoiler     bool
	// Visibility string // TODO: 或许指定用户可见?
	UserID     uint
	User       User
	MetadataID uint
	Metadata   Metadata
}

func (r *Review) Save() (*Review, error) {
	if err := DB.Save(r).Error; err != nil {
		return r, err
	}
	return r, nil
}
