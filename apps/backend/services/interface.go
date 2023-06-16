package services

import "github.com/ryanbradynd05/go-tmdb"

type MovieSourceService interface {
	MovieInfo(id int, lang string) (*tmdb.Movie, error)
	MovieInfoViaAPI(id int, options options) (*tmdb.Movie, error)
	MovieSearch(name string, options options) (*tmdb.MovieSearchResults, error)
}
