package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"travisroad/gotracker/models"

	"github.com/ryanbradynd05/go-tmdb"
	"gorm.io/gorm"
)

type MovieService struct {
	mss MovieSourceService
}

func NewMovieService(ts *TmdbService) *MovieService {
	return &MovieService{
		mss: ts,
	}
}

// GetMovieMetaData retrieves movie metadata using the given ID and source.
// first check whether the movie exists
// if not exist, fetch the movie metadata from the source
// else return the data from database
//
// Parameters:
// - id: the ID of the movie
// - source: the source of the movie
// - options: additional options for retrieving the metadata
//
// Returns:
// - *models.Movie: the retrieved movie metadata
// - error: an error if the retrieval fails
func (ms *MovieService) GetMovieMetaData(id int, source string, options options) (*models.Movie, error) {
	var m *models.Movie = &models.Movie{}
	err := models.DB.Where("identifier = ? and source = ?", id, source).Take(m).Error
	if err == nil {
		return m, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	chEn := make(chan *tmdb.Movie, 1)
	chEnErr := make(chan error, 1)
	chCn := make(chan *tmdb.Movie, 1)
	chCnErr := make(chan error, 1)

	go func() {
		tm, err := ms.mss.MovieInfo(id, "") // default en-US
		chEnErr <- err
		chEn <- tm
	}()
	go func() {
		tm, err := ms.mss.MovieInfo(id, "zh-CN")
		chCnErr <- err
		chCn <- tm
	}()

	tm := <-chEn
	if tm == nil {
		return nil, <-chEnErr
	}
	tmCN := <-chCn
	if tm == nil {
		return nil, <-chCnErr
	}

	t, err := time.Parse("2006-01-02", tm.ReleaseDate)
	if err != nil {
		return nil, err
	}

	m = &models.Movie{
		Time:       int(tm.Runtime),
		Source:     source,
		Identifier: strconv.Itoa(tm.ID),
		Metadata: models.Metadata{
			Title:         tm.Title,
			Description:   tm.Overview,
			PublishYear:   strconv.Itoa(t.Year()),
			PublishData:   tm.ReleaseDate,
			ImageUrl:      "http://image.tmdb.org/t/p/original" + tm.PosterPath, // FIXME:dummy
			SourceUrl:     fmt.Sprintf("https://www.themoviedb.org/movie/%d", id),
			TitleCN:       tmCN.Title,
			DescriptionCN: tmCN.Overview,
			ImageUrlCN:    tmCN.PosterPath,
		},
	}
	err = m.Save()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MovieService) Query(q string, options options) (*tmdb.MovieSearchResults, error) {
	return m.mss.MovieSearch(q, options)
}
