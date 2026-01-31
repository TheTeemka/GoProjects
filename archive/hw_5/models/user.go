package models

type UserDTO struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash []byte `json:"-"`
}

type UserEntity struct {
	ID           int
	Email        string
	PasswordHash []byte
}

func (dto *UserDTO) ToUserEntity() *UserEntity {
	return &UserEntity{
		ID:           dto.ID,
		Email:        dto.Email,
		PasswordHash: dto.PasswordHash,
	}
}

func (entity *UserEntity) ToUserDTO() *UserDTO {
	return &UserDTO{
		ID:           entity.ID,
		Email:        entity.Email,
		PasswordHash: entity.PasswordHash,
	}
}

type CreateUserRequest struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	PlainPassword string `json:"password"`
}

type LoginUserRequest struct {
	Email         string `json:"email"`
	PlainPassword string `json:"password"`
}
