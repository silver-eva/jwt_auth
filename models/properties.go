package models

import "github.com/google/uuid"

type Properties struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID         uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	MaxSessions    int       `json:"max_sessions" gorm:"default:5"`
	SessionExpires int       `json:"session_expires" gorm:"default:7"`
}