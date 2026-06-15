package repo

import (
	"chatapp/internal/model"
	"chatapp/internal/request"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var Clients = make(map[string]*websocket.Conn)

func MessageWebSoketConnection(userID string, c *gin.Context) (interface{}, error) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return nil, err
	}
	Clients[userID] = conn
	defer func() {
		log.Println("WebSocket Disconnected")
		delete(Clients, userID)
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read Error:", err)
			break
		}

		request := request.NewMessageRequest()
		if err := json.Unmarshal(message, request); err != nil {
			log.Println("Read Error:", err)
			break
		}

		result, err := model.GetConversationByConversationID(request.ConversationId)
		if err != nil {
			log.Println("Read Error:", err)
			break
		}
		fmt.Println(userID)
		conversationMembers := result.ConversationMember.String()
		// skip sender
		if result.ConversationMember.String() == userID {
			conversationMembers = result.CreatedBy.String()
		}
		receiverConn, ok := Clients[conversationMembers]
		log.Println("Receiver ID:", conversationMembers)
		log.Println("Sender ID:", userID)
		if ok {

			err = receiverConn.WriteJSON(gin.H{
				"type":     "new_message",
				"senderId": userID,
				"message":  request.Content,
			})

			if err != nil {
				log.Println("Write Error:", err)
			}
		}

		conn.WriteJSON(gin.H{
			"type":    "ack",
			"message": "Delivered",
		})
	}
	return nil, nil
}
