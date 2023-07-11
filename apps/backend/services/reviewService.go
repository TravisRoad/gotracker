package services

import "travisroad/gotracker/models"

type ReviewService struct {
}

func NewReviewService() *ReviewService {
	return &ReviewService{}
}

func (rs *ReviewService) SaveReview(review *models.Review) (*models.Review, error) {
	r, err := review.Save()
	if err != nil {
		return nil, err
	}
	return r, nil
}
