package middlewares

import (
	"github.com/gin-gonic/gin"
)

func OwnsResource() gin.HandlerFunc {
	return func(c *gin.Context) {
		// uid := c.GetUint("uid")
		// if !userOwnsResource(uid, resourceID) {
		// 	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		// 		"error": "You do not have permission to access this resource.",
		// 	})
		// 	return
		// }

		// 继续处理请求
		c.Next()
	}
}

// func userOwnsResource(userID string, resourceID string) bool {
// 	// 实现此函数，根据用户ID和资源ID检查用户是否拥有此资源。
// 	// 返回true或false。
// 	return true
// }
