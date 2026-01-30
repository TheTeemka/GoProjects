package services

import (
	"context"
	"time"

	"github.com/TheTeemka/GoProjects/hw_6/utils"
)

type ITokenRepository interface {
	CreateToken(ctx context.Context, userID int, token string, expiryDate time.Time, tokenType string) error
	TokenExists(ctx context.Context, token string, tokenType string) (int, error)
}

type TokenService struct {
	tokenRepo ITokenRepository
	tokenSize int
	tokenTTL  time.Duration
	tokenType string
}

func NewTokenService(tokenRepo ITokenRepository, tokenSize int, tokenType string, tokenTTL time.Duration) *TokenService {
	return &TokenService{
		tokenRepo: tokenRepo,
		tokenSize: tokenSize,
		tokenType: tokenType,
		tokenTTL:  tokenTTL,
	}
}

func (ts *TokenService) CreateToken(userID int) (string, error) {
	tokenSize := ts.tokenSize
	if tokenSize <= 0 {
		tokenSize = 32
	}

	token, err := utils.GenerateToken(tokenSize)
	if err != nil {
		return "", err
	}
	expiry := time.Now().Add(ts.tokenTTL)
	return token, ts.tokenRepo.CreateToken(context.Background(), userID, token, expiry, ts.tokenType)
}

func (ts *TokenService) ValidateToken(token string) (int, error) {
	return ts.tokenRepo.TokenExists(context.Background(), token, ts.tokenType)
}
