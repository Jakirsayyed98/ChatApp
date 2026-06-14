package middleware

import (
	"chatapp/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "Missing authorization header"})
			c.Abort()
			return
		}

		if len(strings.Split(tokenString, " ")) != 2 {
			c.JSON(401, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		if tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		userID, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("userID", userID.UserID)
	}
}
