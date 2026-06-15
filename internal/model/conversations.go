package model

import (
	"chatapp/internal/config"
	"context"

	"github.com/google/uuid"
)

type Conversations struct {
	ID                 *uuid.UUID `json:"id" db:"id"`
	ConversationType   string     `json:"type" validate:"required,oneof='one-to-one group'"`
	Name               string     `json:"name"`
	ConversationMember uuid.UUID  `json:"conversation_members"`
	CreatedBy          uuid.UUID  `json:"created_by"`
	CreatedAt          string     `json:"created_at"`
	UpdatedAt          string     `json:"updated_at"`
}

type ConversationsDetail struct {
	ID                 uuid.UUID `json:"id"`
	ConversationType   string    `json:"type" validate:"required,oneof='one-to-one group'"`
	Name               string    `json:"name"`
	ConversationMember User      `json:"conversation_members"`
	CreatedBy          User      `json:"created_by"`
	CreatedAt          string    `json:"created_at"`
	UpdatedAt          string    `json:"updated_at"`
}

func InsertConversation(conversation Conversations) error {
	db := config.GetDB()
	txn, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer txn.Rollback()
	query := "INSERT INTO conversations (type, name,conversation_members, created_by) VALUES ($1, $2, $3, $4)"
	_, err = txn.Exec(query, conversation.ConversationType, conversation.Name, conversation.ConversationMember, conversation.CreatedBy)
	if err != nil {
		return err
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	return nil
}

func GetConversationByUserID(userId string) ([]Conversations, error) {
	db := config.GetDB()
	query := "SELECT id, type, name, created_by, created_at, updated_at from conversations where created_by=$1"

	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var conversations []Conversations
	for rows.Next() {
		var conversation Conversations
		if err := rows.Scan(
			&conversation.ID,
			&conversation.ConversationType,
			&conversation.Name,
			&conversation.CreatedBy,
			&conversation.CreatedAt,
			&conversation.UpdatedAt,
		); err != nil {
			return nil, err
		}

		conversations = append(conversations, conversation)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return conversations, err
}

func GetConversationByUserIDAndConversationID(userId, conversationId string) (*ConversationsDetail, error) {

	db := config.GetDB()
	query := "select u.id, u.name, u.email, u.status, u2.id, u2.name, u2.email, u2.status, c.id, c.type, c.name, c.created_at, c.updated_at  from conversations c left join users u on created_by= u.id left join users u2 on conversation_members=u2.id where c.created_by=$1 AND c.id =$2"
	var conversation ConversationsDetail
	err := db.QueryRow(query, userId, conversationId).Scan(
		&conversation.CreatedBy.ID,
		&conversation.CreatedBy.Name,
		&conversation.CreatedBy.Email,
		&conversation.CreatedBy.Status,
		&conversation.ConversationMember.ID,
		&conversation.ConversationMember.Name,
		&conversation.ConversationMember.Email,
		&conversation.ConversationMember.Status,
		&conversation.ID,
		&conversation.ConversationType,
		&conversation.Name,
		&conversation.CreatedAt,
		&conversation.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &conversation, err
}

func GetConversationByConversationID(conversationId string) (*Conversations, error) {

	db := config.GetDB()
	query := "SELECT id, type, name,conversation_members, created_by, created_at, updated_at from conversations where id=$1"
	var conversation Conversations
	if err := db.QueryRow(query, conversationId).Scan(
		&conversation.ID,
		&conversation.ConversationType,
		&conversation.Name,
		&conversation.ConversationMember,
		&conversation.CreatedBy,
		&conversation.CreatedAt,
		&conversation.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &conversation, nil
}
