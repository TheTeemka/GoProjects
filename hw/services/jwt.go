package services

import (
	"time"

	"github.com/TheTeemka/GoProjects/hw_6/models"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey []byte
	ttl       time.Duration
}

func NewJWTService(secretKey []byte, ttl time.Duration) *JWTService {
	return &JWTService{
		secretKey: secretKey,
		ttl:       ttl,
	}
}

func (js *JWTService) CreateToken(user *models.UserDTO) (string, error) {
	claims := &models.UserClaims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "user_claims",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(js.ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString((js.secretKey))
}

func (js *JWTService) ParseToken(tokenStr string) (*models.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(js.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}
