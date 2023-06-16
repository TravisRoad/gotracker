package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Time       int
	MetadataID uint
	Metadata   Metadata
}

func (m *Movie) Save() (*Movie, error) {
	if err := DB.Save(m).Error; err != nil {
		return m, err
	}
	return m, nil
}
