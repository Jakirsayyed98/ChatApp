package repo

import (
	"chatapp/internal/model"
	"chatapp/internal/request"
	"chatapp/internal/response"
	"errors"

	"github.com/google/uuid"
)

func GetConversationByUserID(userId string) ([]response.ConversationResponse, error) {
	if userId == "" {
		return nil, errors.New("Invalid User")
	}

	result, err := model.GetConversationByUserID(userId)
	if err != nil {
		return nil, err
	}

	var conversations []response.ConversationResponse
	for _, v := range result {
		var conversation response.ConversationResponse
		conversation.ID = v.ID
		conversation.ConversationType = v.ConversationType
		conversation.Name = v.Name
		conversation.CreatedBy = v.CreatedBy
		conversation.CreatedAt = v.CreatedAt
		conversation.UpdatedAt = v.UpdatedAt

		conversations = append(conversations, conversation)
	}
	return conversations, nil
}

func CreateUserConversations(userId string, request *request.RequestConversation) error {

	user, err := model.GetUserByMail(request.Email)
	if err != nil {
		return errors.New("User not found")
	}
	userIDCon, err := uuid.Parse(userId)
	if err != nil {
		return err
	}
	conversation := model.Conversations{
		ConversationType:   request.ConversationType,
		Name:               request.Name,
		ConversationMember: user.ID,
		CreatedBy:          userIDCon,
	}

	if err := model.InsertConversation(conversation); err != nil {
		return err
	}
	return nil
}

func GetConversationDetail(userID, conversationId string) (*model.ConversationsDetail, error) {
	if userID == "" || conversationId == "" {
		return nil, errors.New("please select correct conversation")
	}

	result, err := model.GetConversationByUserIDAndConversationID(userID, conversationId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
