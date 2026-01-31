package services

import (
	"context"
	"fmt"

	"github.com/TheTeemka/GoProjects/hw/errs"
	"github.com/TheTeemka/GoProjects/hw/models"
	"github.com/TheTeemka/GoProjects/hw/utils"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *models.UserEntity) error
	GetUserByEmail(ctx context.Context, email string) (*models.UserEntity, error)
	GetUserByID(ctx context.Context, user_id int) (*models.UserEntity, error)
}

type IJwtService interface {
	CreateToken(user *models.UserDTO) (string, error)
}

type ITokenService interface {
	CreateToken(userID int) (string, error)
	ValidateToken(token string) (int, error)
}

type AuthService struct {
	userRepo            IUserRepository
	jwtService          IJwtService
	refreshTokenService ITokenService
}

func NewAuthService(userRepo IUserRepository, jwtService IJwtService,
	refreshTokenService ITokenService) *AuthService {
	return &AuthService{
		userRepo:            userRepo,
		jwtService:          jwtService,
		refreshTokenService: refreshTokenService,
	}
}

func (s *AuthService) CreateUser(dto *models.CreateUserRequest) error {
	passwordHash, err := utils.HashPassword([]byte(dto.PlainPassword))
	if err != nil {
		return fmt.Errorf("err in CreateUser: %w", err)
	}

	userRole, err := models.ParseUserRole(dto.Role)
	if err != nil {
		return fmt.Errorf("err in CreateUser: %w", err)
	}

	entity := &models.UserEntity{
		Email:        dto.Email,
		Role:         userRole,
		PasswordHash: passwordHash,
	}

	return s.userRepo.CreateUser(context.Background(), entity)
}

func (s *AuthService) GetUserByEmail(email string) (*models.UserDTO, error) {
	userEntity, err := s.userRepo.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, fmt.Errorf("err in GetUserByEmail: %w", err)
	}

	return userEntity.ToUserDTO(), nil
}

func (s *AuthService) Login(email, plainPassword string) (accessToken string, refreshToken string, err error) {
	userEntity, err := s.userRepo.GetUserByEmail(context.Background(), email)
	if err != nil {
		return "", "", fmt.Errorf("err in Login: %w", err)
	}

	if userEntity == nil {
		return "", "", errs.ErrUserNotFound
	}

	if err := utils.ComparePassword(userEntity.PasswordHash, []byte(plainPassword)); err != nil {
		return "", "", errs.ErrPasswordMismatch
	}

	accessToken, err = s.jwtService.CreateToken(userEntity.ToUserDTO())
	if err != nil {
		return "", "", fmt.Errorf("err in Login: %w", err)
	}

	refreshToken, err = s.refreshTokenService.CreateToken(userEntity.ID)
	if err != nil {
		return "", "", fmt.Errorf("err in Login: %w", err)
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) RefreshAccessToken(refreshToken string) (string, error) {
	user_id, err := s.refreshTokenService.ValidateToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("err in RefreshAccessToken: %w", err)
	}

	userEntity, err := s.userRepo.GetUserByID(context.Background(), user_id)
	if err != nil {
		return "", fmt.Errorf("err in RefreshAccessToken: %w", err)
	}

	accessToken, err := s.jwtService.CreateToken(userEntity.ToUserDTO())
	if err != nil {
		return "", fmt.Errorf("err in RefreshAccessToken: %w", err)
	}

	return accessToken, nil
}
