package handler

import (
	"chatapp/internal/handler/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) {
	authGroup := app.Group("/api")
	auth.AuthRouter(authGroup)
	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "healthy",
		})
	})
}
