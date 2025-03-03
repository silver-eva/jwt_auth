package models

type RegisterRequest struct {
	Uname string `json:"uname"`
	Name     string `json:"name"`
	Second   string `json:"second"`
	Password string `json:"password"`
}