package response

type LoginUserResponse struct {
	Token string `json:"token"`
}

type CurrentUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}
