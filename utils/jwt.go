package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)


type JWTUtils struct {
	secret string
	expires int
}

func (j *JWTUtils) GenerateAccessToken(userID uuid.UUID) (string, error) {
	claims :=  jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Duration(j.expires) * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

func (j *JWTUtils) GenerateRefreshToken(userID uuid.UUID, days int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Duration(days) * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

func (j *JWTUtils) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

var JWT *JWTUtils

func InitJWT(secret string, expiries int) {
	JWT = &JWTUtils{
		secret: secret,
		expires: expiries,
	}
}

