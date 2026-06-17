package repo

import (
	"chatapp/internal/model"
	"chatapp/internal/request"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

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
	conn.SetReadDeadline(time.Now().Add(1 * time.Minute))
	conn.SetPongHandler(func(appData string) error {
		conn.SetReadDeadline(time.Now().Add(1 * time.Minute))
		return nil
	})

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
			if errors.Is(err, sql.ErrNoRows) {
				log.Println("Conversation not found:", request.ConversationId)

				conn.WriteJSON(gin.H{
					"type":  "error",
					"error": "conversation not found",
				})

				continue
			}
			log.Println("Read Error:", err)
			break
		}

		conversationMembers := result.ConversationMember.String()
		// skip sender
		if result.ConversationMember.String() == userID {
			conversationMembers = result.CreatedBy.String()
		}

		receiverConn, ok := Clients[conversationMembers]
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
	}
	return nil, nil
}
