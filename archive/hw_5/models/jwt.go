package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID int
	Email  string
	jwt.RegisteredClaims
}
