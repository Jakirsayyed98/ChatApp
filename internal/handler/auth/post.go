package auth

import (
	"chatapp/internal/repo"
	"chatapp/internal/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	request := request.NewRegisterUserRequest()
	if err := request.Bind(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := repo.RegisterUser(request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})

}

func LoginUser(c *gin.Context) {
	request := request.NewLoginUserRequest()
	if err := request.Bind(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := repo.LoginUser(request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": response.Token})
}
