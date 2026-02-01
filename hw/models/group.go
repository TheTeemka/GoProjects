package models

import "time"

type GroupDTO struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	EnrollmentYear int        `json:"enrollment_year"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

type CreateGroupRequest struct {
	Name           string `json:"name" validate:"required"`
	EnrollmentYear int    `json:"enrollment_year" validate:"required"`
}

type GroupEntity struct {
	ID             int
	Name           string
	EnrollmentYear int
	CreatedAt      time.Time
	UpdatedAt      *time.Time
}

func (entity *GroupEntity) ToDTO() *GroupDTO {
	if entity == nil {
		return nil
	}
	return &GroupDTO{
		ID:             entity.ID,
		Name:           entity.Name,
		EnrollmentYear: entity.EnrollmentYear,
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
	}
}

func (req *CreateGroupRequest) ToEntity() *GroupEntity {
	return &GroupEntity{
		Name:           req.Name,
		EnrollmentYear: req.EnrollmentYear,
	}
}
