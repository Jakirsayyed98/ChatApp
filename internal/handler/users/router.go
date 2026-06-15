package users

import (
	"chatapp/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(app *gin.RouterGroup) {
	auth := app.Group("/users")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/me", GetCurrentUser)
}
