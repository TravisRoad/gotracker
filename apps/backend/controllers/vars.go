package controllers

import (
	"travisroad/gotracker/auth"
	"travisroad/gotracker/services"
)

var (
	jwtHelper     *auth.JWTAuthHelper
	movieService  *services.MovieService
	reviewService *services.ReviewService
)

func Inject(jh *auth.JWTAuthHelper, ms *services.MovieService, rs *services.ReviewService) {
	jwtHelper = jh
	movieService = ms
	reviewService = rs
}

type AddMovieReviewInput struct {
	Id      string   `json:"id" binding:"required"`
	Source  string   `json:"source" binding:"required"`
	Rank    []string `json:"rank" binding:"required"` // element is [0, 20] string
	Content string   `json:"content" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
