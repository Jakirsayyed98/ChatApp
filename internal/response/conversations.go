package response

import "github.com/google/uuid"

type ConversationResponse struct {
	ID               *uuid.UUID `json:"id"`
	ConversationType string     `json:"type"`
	Name             string     `json:"name"`
	CreatedBy        uuid.UUID  `json:"created_by"`
	CreatedAt        string     `json:"created_at"`
	UpdatedAt        string     `json:"updated_at"`
}
