package messages

import (
	"chatapp/internal/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMessagesHandler(c *gin.Context) {
	userId, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id not found"})
		return
	}

	_, err := repo.MessageWebSoketConnection(userId.(string), c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
