package request

import "github.com/gin-gonic/gin"

type MessageRequest struct {
	ConversationId string `json:"conversationId"`
	MessageType    string `json:"messageType"`
	Content        string `json:"content"`
	FileUrl        string `json:"fileUrl"`
	ReplyToID      string `json:"replyToId"`
}

func NewMessageRequest() *MessageRequest {
	return &MessageRequest{}
}

func (r *MessageRequest) Bind(c *gin.Context) error {
	return c.ShouldBind(r)
}
