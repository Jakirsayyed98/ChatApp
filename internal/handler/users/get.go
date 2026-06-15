package users

import (
	"chatapp/internal/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "User not authenticated"})
		return
	}

	result, err := repo.GetCurrentUserByID(userID.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result, "message": "User fetched successfully"})
}
