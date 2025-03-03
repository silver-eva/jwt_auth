package models

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Created    time.Time `json:"created" gorm:"default:now()"`
	Refreshed  time.Time `json:"refreshed" gorm:"default:now()"`
	Terminated *time.Time `json:"terminated"`
	Token    string    	`json:"token" gorm:"not null"`
	Active     bool      `json:"active" gorm:"default:true"`
}