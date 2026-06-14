package request

import "github.com/gin-gonic/gin"

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func NewRegisterUserRequest() *RegisterUserRequest {
	return &RegisterUserRequest{}
}

func (r *RegisterUserRequest) Bind(c *gin.Context) error {
	return c.ShouldBind(r)
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewLoginUserRequest() *LoginUserRequest {
	return &LoginUserRequest{}
}

func (r *LoginUserRequest) Bind(c *gin.Context) error {
	return c.ShouldBind(r)
}
