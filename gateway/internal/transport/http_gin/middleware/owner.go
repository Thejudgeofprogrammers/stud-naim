package middleware

import "github.com/gin-gonic/gin"

func OwnerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("user_id")
		paramID := c.Param("id")

		if userID == "" || paramID == "" || userID != paramID {
			c.AbortWithStatusJSON(403, gin.H{"error": "forbidden"})
			return
		}

		c.Next()
	}
}
