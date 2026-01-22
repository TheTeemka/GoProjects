package services

import (
	"context"
	"fmt"

	"github.com/TheTeemka/GoProjects/hw_5/models"
	"github.com/TheTeemka/GoProjects/hw_5/utils"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *models.UserEntity) error
	GetUserByEmail(ctx context.Context, email string) (*models.UserEntity, error)
}

type UserService struct {
	userRepo IUserRepository
}

func NewUserService(userRepo IUserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) CreateUser(dto *models.CreateUserRequest) error {
	passwordHash, err := utils.HashPassword([]byte(dto.PlainPassword))
	if err != nil {
		return fmt.Errorf("err in CreateUser: %w", err)
	}

	entity := &models.UserEntity{
		Email:        dto.Email,
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
