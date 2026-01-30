package services

import (
	"context"
	"fmt"

	"github.com/TheTeemka/GoProjects/hw_6/models"
	"github.com/TheTeemka/GoProjects/hw_6/utils"
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

type UserService struct {
	userRepo            IUserRepository
	jwtService          IJwtService
	refreshTokenService ITokenService
}

func NewUserService(userRepo IUserRepository, jwtService IJwtService,
	refreshTokenService ITokenService) *UserService {
	return &UserService{
		userRepo:            userRepo,
		jwtService:          jwtService,
		refreshTokenService: refreshTokenService,
	}
}

func (us *UserService) CreateUser(dto *models.CreateUserRequest) error {
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

	return us.userRepo.CreateUser(context.Background(), entity)
}

func (us *UserService) GetUserByEmail(email string) (*models.UserDTO, error) {
	userEntity, err := us.userRepo.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, fmt.Errorf("err in GetUserByEmail: %w", err)
	}

	return userEntity.ToUserDTO(), nil
}

func (us *UserService) Login(email, plainPassword string) (accessToken string, refreshToken string, err error) {
	userEntity, err := us.userRepo.GetUserByEmail(context.Background(), email)
	if err != nil {
		return "", "", fmt.Errorf("err in Login: %w", err)
	}

	if userEntity == nil {
		return "", "", models.ErrUserNotFound
	}

	if err := utils.ComparePassword(userEntity.PasswordHash, []byte(plainPassword)); err != nil {
		return "", "", models.ErrPasswordMismatch
	}

	accessToken, err = us.jwtService.CreateToken(userEntity.ToUserDTO())
	if err != nil {
		return "", "", fmt.Errorf("err in Login: %w", err)
	}

	refreshToken, err = us.refreshTokenService.CreateToken(userEntity.ID)
	if err != nil {
		return "", "", fmt.Errorf("err in Login: %w", err)
	}

	return accessToken, refreshToken, nil
}

func (us *UserService) RefreshAccessToken(refreshToken string) (string, error) {
	user_id, err := us.refreshTokenService.ValidateToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("err in RefreshAccessToken: %w", err)
	}

	userEntity, err := us.userRepo.GetUserByID(context.Background(), user_id)
	if err != nil {
		return "", fmt.Errorf("err in RefreshAccessToken: %w", err)
	}

	accessToken, err := us.jwtService.CreateToken(userEntity.ToUserDTO())
	if err != nil {
		return "", fmt.Errorf("err in RefreshAccessToken: %w", err)
	}

	return accessToken, nil
}
