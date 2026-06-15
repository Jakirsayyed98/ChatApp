package conversation

import (
	"chatapp/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ConversationRouter(app *gin.RouterGroup) {
	app.Use(middleware.AuthMiddleware())
	app.GET("conversations", GetConversationByUserID)
	app.POST("conversations", CreateConversations)
	app.GET("conversations/:id", GetConversationDetailByID)
}
