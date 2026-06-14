package auth

import (
	"chatapp/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRouter(app *gin.RouterGroup) {
	auth := app.Group("/auth")
	auth.POST("/register", RegisterUser)
	auth.POST("/login", LoginUser)

	auth.Use(middleware.AuthMiddleware())
	auth.GET("/current-user", GetCurrentUser)
}
