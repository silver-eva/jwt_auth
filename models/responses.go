package models

type ErrorResponse struct {
	Message string `json:"message"`
}

type LoggedInUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}