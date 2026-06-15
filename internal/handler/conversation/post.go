package conversation

import (
	"chatapp/internal/repo"
	"chatapp/internal/request"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateConversations(c *gin.Context) {
	userId, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID Not found"})
		return
	}

	conversation := request.NewRequestConversation()

	if err := conversation.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := repo.CreateUserConversations(userId.(string), conversation)
	if err != nil {
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Conversation Created Successfully"})
}
