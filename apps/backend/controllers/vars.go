package controllers

import (
	"travisroad/gotracker/auth"
	"travisroad/gotracker/services"
)

var (
	jwtHelper    *auth.JWTAuthHelper
	movieService *services.MovieService
)

func Inject(jh *auth.JWTAuthHelper, ms *services.MovieService) {
	jwtHelper = jh
	movieService = ms
}
