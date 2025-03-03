package models

type RegisterRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Second   string `json:"second"`
	Password string `json:"password"`
}