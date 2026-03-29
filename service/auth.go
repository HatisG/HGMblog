package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	SecretKey []byte
}

func (s *AuthService) GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(s.SecretKey)
}

func (s *AuthService) VerifyToken(tokenSting string) (uint, error) {
	token, err := jwt.Parse(tokenSting, func(t *jwt.Token) (any, error) {
		return s.SecretKey, nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("无效token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("无法解析token")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("无法获取ID")
	}
	return uint(userID), nil

}
