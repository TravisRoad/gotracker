package services

import "travisroad/gotracker/models"

type SeenService struct{}

func (s *SeenService) IsExistMovieSeen(uid uint, identifier string, source string) (bool, error) {
	var cnt int64
	if err := models.DB.Model(&models.Seen{}).Where("uid = ? and identifier = ? and source = ? and variety = ?", uid, identifier, source, "movie").Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}
