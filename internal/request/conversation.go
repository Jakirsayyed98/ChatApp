package request

import "github.com/gin-gonic/gin"

type RequestConversation struct {
	ConversationType string `json:"type" validate:"required,oneof='one-to-one group'"`
	Name             string `json:"name"`
	Email            string `json:"email"`
}

func NewRequestConversation() *RequestConversation {
	return &RequestConversation{}
}

func (request *RequestConversation) Bind(c *gin.Context) error {
	return c.ShouldBind(request)
}
