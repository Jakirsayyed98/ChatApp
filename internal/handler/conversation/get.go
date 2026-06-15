package conversation

import (
	"chatapp/internal/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetConversationByUserID(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid User",
		})
	}

	result, err := repo.GetConversationByUserID(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": "Successfully get conversations",
	})
}

func GetConversationDetailByID(c *gin.Context) {
	conversationID := c.Param("id")
	userId, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user id not found",
		})
		return
	}

	result, err := repo.GetConversationDetail(userId.(string), conversationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result, "message": "Successfully get Conversation Details"})
}
