package route

import (
	"travisroad/gotracker/controllers"
	"travisroad/gotracker/middlewares"

	"github.com/gin-gonic/gin"
)

func RouteInit() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.Register)

	search := r.Group("/api/search", middlewares.JwtAuthMiddleware())
	search.GET("/movie", controllers.Search)

	return r
}
