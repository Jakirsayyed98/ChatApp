package model

import (
	"chatapp/internal/config"
	"context"

	"github.com/google/uuid"
)

type Messages struct {
	ID             uuid.UUID `json:"id" db:"id"`
	ConversationId uuid.UUID `json:"conversation_id" db:"conversation_id"`
	SenderId       uuid.UUID `json:"sender_id" db:"sender_id"`
	MessageType    string    `json:"message_type" db:"message_type"`
	Content        string    `json:"content" db:"content"`
	FileUrl        string    `json:"file_url" db:"file_url"`
	ReplyToID      uuid.UUID `json:"reply_to_id" db:"reply_to_id"`
	CreatedAt      string    `json:"created_at" db:"created_at"`
	UpdatedAt      string    `json:"updated_at" db:"updated_at"`
}

func InsertMessages(message Messages) error {
	db := config.GetDB()
	query := "INSERT INTO messages (conversation_id, sender_id, message_type, content, file_url, reply_to_id) VALUES ($1, $2, $3, $4, $5, $6)"
	txn, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer txn.Rollback()
	_, err = txn.Exec(query, message.ConversationId, message.SenderId, message.MessageType, message.Content, message.FileUrl, message.ReplyToID)
	if err != nil {
		return err
	}
	if err := txn.Commit(); err != nil {
		return err
	}
	return nil
}
