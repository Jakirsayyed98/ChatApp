package repo

import (
	"chatapp/internal/model"
	"chatapp/internal/request"
	"chatapp/internal/response"
	"chatapp/internal/utils"
	"fmt"
)

func RegisterUser(request *request.RegisterUserRequest) error {
	if request.Email == "" || request.Password == "" || request.Username == "" {
		return fmt.Errorf("email, password and username are required")
	}

	md5Hash := fmt.Sprintf("%x", request.Password)

	user := &model.User{
		Username: request.Username,
		Password: md5Hash,
		Email:    request.Email,
	}
	return model.InsertUser(user)
}

func LoginUser(request *request.LoginUserRequest) (*response.LoginUserResponse, error) {
	userData, err := model.GetUserByMail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("user not found", err)
	}

	if userData.Password != fmt.Sprintf("%x", request.Password) {
		return nil, fmt.Errorf("invalid password")
	}

	token, err := utils.CreateToken(string(userData.ID.String()))
	if err != nil {
		return nil, fmt.Errorf("failed to create token")
	}

	return &response.LoginUserResponse{
		Token: token, // Replace this with actual JWT token generation
	}, nil
}

func GetCurrentUserByID(userID string) (*response.CurrentUserResponse, error) {
	if userID == "" {
		return nil, fmt.Errorf("user ID is required")
	}

	user, err := model.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	userResponse := &response.CurrentUserResponse{
		ID:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
		Status:   user.Status,
	}
	return userResponse, nil
}
