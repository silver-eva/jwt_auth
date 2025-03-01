package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username string    `json:"username" gorm:"unique;not null"`
	Name     string    `json:"name"`
	Second   string    `json:"second"`
	Password string    `json:"-" gorm:"not null"`
	Sessions []Session `json:"sessions" gorm:"foreignKey:UserID"`
	Properties Properties `json:"properties" gorm:"foreignKey:UserID"`
	Accesses  []Access   `json:"accesses" gorm:"foreignKey:UserID"`
}