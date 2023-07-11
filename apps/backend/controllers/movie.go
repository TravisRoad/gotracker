package controllers

import (
	"bytes"
	"encoding/gob"
	"net/http"
	"strconv"
	"travisroad/gotracker/constants"
	"travisroad/gotracker/models"

	"github.com/gin-gonic/gin"
)

// GET /api/search/movie?q=&lang=
// lang: en_US zh_CN
func Search(c *gin.Context) {
	query := c.Query("q")
	lang := c.Query("lang")
	if len(query) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "query is required"})
		return
	}

	r, err := movieService.Query(query, map[string]string{"language": lang})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}

// AddMovie is a Go function that adds a movie.
//
// c: A pointer to a gin.Context object.
// The function does not return anything.
func AddMovieReview(c *gin.Context) {
	var input AddMovieReviewInput

	// get user id
	uidValue, ok := c.Get("uid")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse(constants.UNAUTHORIZED, "uid not passed in"))
		return
	}
	uid, ok := uidValue.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse(constants.UNAUTHORIZED, "uid not passed in"))
		return
	}

	// get body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(constants.BODYERROR, err.Error()))
		return
	}
	id, err := strconv.Atoi(input.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if user has not seen the movie, return an error
	exist, err := seenService.IsExistMovieSeen(uid, input.Id, input.Source)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": constants.NOTSEEN, "error": "user has not seen the movie"})
		return
	}

	// get movie metadata, save the metadata into database if it does not exist
	m, err := movieService.GetMovieMetaData(id, input.Source, map[string]string{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// compute the final rating
	var rating float32 = 0.0
	for _, v := range input.Rank {
		x, err := strconv.Atoi(v)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		rating += float32(x)
	}
	float32Rating := rating / float32(len(input.Rank)) / 2.0

	// use gob to encode the rank array
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(input.Rank)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "gob: " + err.Error()})
		return
	}

	review := &models.Review{
		Rating:      float32Rating,
		ExtraRating: buf.Bytes(),
		Content:     input.Content,
		MetadataID:  m.MetadataID,
		UserID:      uid,
	}

	_, err = review.Save()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func GetMovieInfo(c *gin.Context) {

}
