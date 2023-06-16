package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /api/search/movie?q=&lang=
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
