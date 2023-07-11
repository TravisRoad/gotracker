package controllers

import (
	"net/http"
	"travisroad/gotracker/models"

	"github.com/gin-gonic/gin"
)

// GetAllReviewsByUser gets all reviews by user.
//
// It takes a gin.Context parameter.
// It does not return anything.
func GetAllReviewsByUser(c *gin.Context) {
	uidValue, ok := c.Get("uid")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "uid not passed"})
		return
	}
	uid, ok := uidValue.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}

	c.Query("page")

	reviews := &[]models.Review{}

	if err := models.DB.Scopes(models.Paginate(c)).Where("user_id = ?", uid).Find(reviews).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
