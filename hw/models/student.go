package models

import "time"

type StudentDTO struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Birthday  time.Time  `json:"birthday"`
	GroupID   int        `json:"group_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type CreateStudentRequest struct {
	Name     string    `json:"name" validate:"required"`
	Email    string    `json:"email" validate:"required,email"`
	Birthday time.Time `json:"birthday" validate:"required"`
	GroupID  int       `json:"group_id" validate:"required"`
}

type UpdateStudentRequest struct {
	Name     *string    `json:"name,omitempty"`
	Email    *string    `json:"email,omitempty" validate:"omitempty,email"`
	Birthday *time.Time `json:"birthday,omitempty"`
	GroupID  *int       `json:"group_id,omitempty"`
}

type StudentEntity struct {
	ID        int
	Name      string
	Email     string
	Birthday  time.Time
	GroupID   int
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type StudentFilter struct {
	Name    *string
	Email   *string
	GroupID *int
	Limit   uint64
	Offset  uint64
}

func (entity *StudentEntity) ToDTO() *StudentDTO {
	if entity == nil {
		return nil
	}
	return &StudentDTO{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Birthday:  entity.Birthday,
		GroupID:   entity.GroupID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func (entity *StudentEntity) PatchFromRequest(req *UpdateStudentRequest) {
	if req.Name != nil {
		entity.Name = *req.Name
	}
	if req.Email != nil {
		entity.Email = *req.Email
	}
	if req.Birthday != nil {
		entity.Birthday = *req.Birthday
	}
	if req.GroupID != nil {
		entity.GroupID = *req.GroupID
	}
}

func (req *CreateStudentRequest) ToEntity() *StudentEntity {
	return &StudentEntity{
		Name:     req.Name,
		Email:    req.Email,
		Birthday: req.Birthday,
		GroupID:  req.GroupID,
	}
}
