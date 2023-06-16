package middlewares

import (
	"net/http"
	"travisroad/gotracker/auth"
	"travisroad/gotracker/di"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		di.C.Invoke(
			func(jh *auth.JWTAuthHelper) {
				uid, err := jh.ExtractTokenID(c)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
					return
				}
				c.Set("uid", uid)
				c.Next()
			})
	}
}
