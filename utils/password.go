package utils

import (
	"jwt_auth/config"
)

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

var PWD PasswordHasher

func InitPassword(cfg *config.Config) {
	switch cfg.Hash {
	case "bcypt":
		PWD = NewBcryptHasher(cfg.HashComplixity)
	case "argon2":
		PWD = NewArgon2Hasher(cfg.HashComplixity)
	default:
		PWD = NewBcryptHasher(cfg.HashComplixity) // default to bcrypt
	}
}
