package messages

import (
	"chatapp/internal/middleware"

	"github.com/gin-gonic/gin"
)

func MessageRouter(app *gin.RouterGroup) {
	app.Use(middleware.AuthMiddleware())
	app.GET("messages", SendMessagesHandler)
}
