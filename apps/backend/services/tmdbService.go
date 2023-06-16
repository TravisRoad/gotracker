package services

import (
	"fmt"
	"travisroad/gotracker/config"
	"travisroad/gotracker/kvstore"
	"travisroad/gotracker/kvstore/lrumap"

	"github.com/ryanbradynd05/go-tmdb"
)

type options map[string]string

type TmdbService struct {
	tmdbAPI   *tmdb.TMDb
	tmdbCache kvstore.KVStore // {[key: string]: *tmdb.Movie}
}

func NewTmdbService(cf *config.Config) *TmdbService {
	c := tmdb.Config{
		APIKey: cf.TmdbKey,
		// APIKey:   "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI0YjZlODIzZWVjMmRhZDlmYmJjY2RiMDgyOWZlOTk5ZCIsInN1YiI6IjY0ODA1M2I0ZTI3MjYwMDEwNzIwMjE4YSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.jVHCmo1qQ9PR2u4bpIDmHYj0q1nfhEJNn7dC4D6y9h8",
		Proxies:  nil,
		UseProxy: false, // TODO: add proxy
	}
	t := &TmdbService{
		tmdbAPI:   tmdb.Init(c),
		tmdbCache: lrumap.New(1_000),
	}
	return t
}

func (t *TmdbService) MovieInfo(id int, lang string) (*tmdb.Movie, error) {
	value, ok := t.tmdbCache.Get(fmt.Sprintf("%d%s", id, lang))
	if v, vok := value.(tmdb.Movie); ok && vok {
		return &v, nil
	}
	movie, err := t.MovieInfoViaAPI(id, map[string]string{
		"language": lang,
	})
	if err != nil {
		return nil, err
	}
	t.tmdbCache.Set(fmt.Sprintf("%d%s", id, lang), *movie) // set value
	return movie, err
}

func (t *TmdbService) MovieInfoViaAPI(id int, options options) (*tmdb.Movie, error) {
	movie, err := t.tmdbAPI.GetMovieInfo(id, options)
	return movie, err
}

func (t *TmdbService) MovieSearch(name string, options options) (*tmdb.MovieSearchResults, error) {
	result, err := t.tmdbAPI.SearchMovie(name, options)
	return result, err
}

func (t *TmdbService) Configuration() (*tmdb.Configuration, error) {
	config, err := t.tmdbAPI.GetConfiguration()
	return config, err
}
