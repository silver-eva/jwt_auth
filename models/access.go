package models

import "github.com/google/uuid"

type Access struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Endpoint string    `json:"endpoint" gorm:"not null"`
}