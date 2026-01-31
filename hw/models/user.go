package models

import (
	"strings"

	"github.com/TheTeemka/GoProjects/hw/errs"
)

type UserDTO struct {
	ID           int      `json:"id"`
	Email        string   `json:"email"`
	Role         UserRole `json:"role"`
	PasswordHash []byte   `json:"-"`
}

type UserEntity struct {
	ID           int
	Email        string
	Role         UserRole
	PasswordHash []byte
}

func (dto *UserDTO) ToUserEntity() *UserEntity {
	return &UserEntity{
		ID:           dto.ID,
		Email:        dto.Email,
		Role:         dto.Role,
		PasswordHash: dto.PasswordHash,
	}
}

func (entity *UserEntity) ToUserDTO() *UserDTO {
	return &UserDTO{
		ID:           entity.ID,
		Email:        entity.Email,
		Role:         entity.Role,
		PasswordHash: entity.PasswordHash,
	}
}

type CreateUserRequest struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	PlainPassword string `json:"password"`
}

type LoginUserRequest struct {
	Email         string `json:"email"`
	PlainPassword string `json:"password"`
}

type UserRole string

const (
	RoleUser    UserRole = "user"
	RoleTeacher UserRole = "teacher"
	RoleAdmin   UserRole = "admin"
)

var roleMap = map[string]UserRole{
	"user":    RoleUser,
	"teacher": RoleTeacher,
	"admin":   RoleAdmin,
}

func ParseUserRole(role string) (UserRole, error) {
	r, exists := roleMap[strings.ToLower(role)]
	if !exists {
		return "", errs.ErrInvalidUserRole
	}
	return r, nil
}
