package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Time       int
	Source     string // dataSource name, like douban, tmdb etc.
	Identifier string
	MetadataID uint `gorm:"index"`
	Metadata   Metadata
}

func (m *Movie) Save() (*Movie, error) {
	if err := DB.Save(m).Error; err != nil {
		return m, err
	}
	return m, nil
}
