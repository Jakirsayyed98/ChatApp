package handler

import (
	"chatapp/internal/handler/auth"
	"chatapp/internal/handler/conversation"
	"chatapp/internal/handler/messages"
	"chatapp/internal/handler/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) {
	apiGroup := app.Group("/api")
	auth.AuthRouter(apiGroup)
	users.UserRouter(apiGroup)
	conversation.ConversationRouter(apiGroup)
	messages.MessageRouter(apiGroup)
	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "healthy",
		})
	})
}
