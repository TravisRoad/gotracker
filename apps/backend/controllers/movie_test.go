package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"travisroad/gotracker/controllers"
	"travisroad/gotracker/models"
	"travisroad/gotracker/utils"

	"github.com/gin-gonic/gin"
)

func TestAddMovieReview(t *testing.T) {
	utils.PreTest()
	user := &models.User{
		Username: "testaddmoviereview",
		Password: "bar",
	}
	user.Save()

	tests := []struct {
		name string
		body controllers.AddMovieReviewInput
		code int
	}{
		{
			name: "success",
			body: controllers.AddMovieReviewInput{
				Id:      "569094",
				Source:  "tmdb",
				Rank:    []string{"18", "18", "18", "18"},
				Content: "idk",
			},
			code: http.StatusOK,
		},
		{
			name: "id fatal",
			body: controllers.AddMovieReviewInput{
				Id:      "1",
				Source:  "tmdb",
				Rank:    []string{"18", "18", "18", "18"},
				Content: "idk",
			},
			code: http.StatusBadRequest,
		},
		{
			name: "rating fatal",
			body: controllers.AddMovieReviewInput{
				Id:      "569094",
				Source:  "tmdb",
				Rank:    []string{"12weq18", "dfas18", "x18", "18"},
				Content: "idk",
			},
			code: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			bodyBytes, err := json.Marshal(tt.body)
			if err != nil {
				t.Errorf(err.Error())
			}
			c.Request, _ = http.NewRequest("POST", "/api/search/movie", bytes.NewBuffer(bodyBytes))
			c.Request.Header.Set("Content-Type", "application/json")
			c.Set("uid", user.ID)

			controllers.AddMovieReview(c)

			if w.Code != tt.code {
				t.Error("got", w.Code, "want", tt.code)
			}
		})
	}
	utils.PostTest()
}
