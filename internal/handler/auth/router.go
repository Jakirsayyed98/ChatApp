package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRouter(app *gin.RouterGroup) {
	auth := app.Group("/auth")
	auth.POST("/register", RegisterUser)
	auth.POST("/login", LoginUser)
}
