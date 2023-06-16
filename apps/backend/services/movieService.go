package services

import (
	"fmt"
	"strconv"
	"time"
	"travisroad/gotracker/models"

	"github.com/ryanbradynd05/go-tmdb"
)

type MovieService struct {
	mss MovieSourceService
}

func NewMovieService(ts *TmdbService) *MovieService {
	return &MovieService{
		mss: ts,
	}
}

func (ms *MovieService) GetMovieMetaData(id int, options options) error {
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
		return <-chEnErr
	}
	tmCN := <-chCn
	if tm == nil {
		return <-chCnErr
	}

	t, err := time.Parse("2006-01-02", tm.ReleaseDate)
	if err != nil {
		return err
	}

	m := &models.Movie{
		Time: int(tm.Runtime),
		Metadata: models.Metadata{
			Title:         tm.Title,
			Description:   tm.Overview,
			PublishYear:   strconv.Itoa(t.Year()),
			PublishData:   tm.ReleaseDate,
			ImageUrl:      "http://image.tmdb.org/t/p/original" + tm.PosterPath, // FIXME:dummy
			Identifier:    strconv.Itoa(tm.ID),
			SourceUrl:     fmt.Sprintf("https://www.themoviedb.org/movie/%d", id),
			TitleCN:       tmCN.Title,
			DescriptionCN: tmCN.Overview,
			ImageUrlCN:    tmCN.PosterPath,
		},
	}
	m.Save()
	return nil
}

func (m *MovieService) Query(q string, options options) (*tmdb.MovieSearchResults, error) {
	return m.mss.MovieSearch(q, options)
}
