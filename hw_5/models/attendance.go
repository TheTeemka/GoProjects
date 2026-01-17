package models

import (
	"fmt"
	"time"
)

type AttendanceDTO struct {
	ID        int       `json:"id"`
	StudentID int       `json:"student_id"`
	SubjectID int       `json:"subject_id"`
	VisitDate time.Time `json:"visit_date"`
	Visited   bool      `json:"visited"`
}

func (entity *AttendanceEntity) ToAttendanceDTO() *AttendanceDTO {
	return &AttendanceDTO{
		ID:        entity.ID,
		StudentID: entity.StudentID,
		SubjectID: entity.SubjectID,
		VisitDate: entity.VisitDate,
		Visited:   entity.Visited,
	}
}

type AttendanceEntity struct {
	ID        int       `json:"id"`
	StudentID int       `json:"student_id"`
	SubjectID int       `json:"subject_id"`
	VisitDate time.Time `json:"visit_date"`
	Visited   bool      `json:"visited"`
}

func (dto *AttendanceDTO) ToAttendanceEntity() *AttendanceEntity {
	return &AttendanceEntity{
		ID:        dto.ID,
		StudentID: dto.StudentID,
		SubjectID: dto.SubjectID,
		VisitDate: dto.VisitDate,
		Visited:   dto.Visited,
	}
}

type CreateAttendanceRequest struct {
	StudentID int    `json:"student_id"`
	SubjectID int    `json:"subject_id"`
	VisitDate string `json:"visit_date"`
	Visited   bool   `json:"visited"`
}

func (req *CreateAttendanceRequest) ToCreateAttendanceDTO() (*CreateAttendanceDTO, error) {
	visitDate, err := time.Parse("2006-01-02", req.VisitDate)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %v", err)
	}
	return &CreateAttendanceDTO{
		StudentID: req.StudentID,
		SubjectID: req.SubjectID,
		VisitDate: visitDate,
		Visited:   req.Visited,
	}, nil
}

type CreateAttendanceDTO struct {
	StudentID int       `json:"student_id"`
	SubjectID int       `json:"subject_id"`
	VisitDate time.Time `json:"visit_date"`
	Visited   bool      `json:"visited"`
}

func (dto *CreateAttendanceDTO) ToAttendanceEntity() *AttendanceEntity {
	return &AttendanceEntity{
		StudentID: dto.StudentID,
		SubjectID: dto.SubjectID,
		VisitDate: dto.VisitDate,
		Visited:   dto.Visited,
	}
}

type AttendanceResponse struct {
	ID        int    `json:"id"`
	StudentID int    `json:"student_id"`
	SubjectID int    `json:"subject_id"`
	VisitDate string `json:"visit_date"`
	Visited   bool   `json:"visited"`
}

func (dto *AttendanceDTO) ToAttendanceResponse() *AttendanceResponse {
	return &AttendanceResponse{
		ID:        dto.ID,
		StudentID: dto.StudentID,
		SubjectID: dto.SubjectID,
		VisitDate: dto.VisitDate.Format("2006-01-02"),
		Visited:   dto.Visited,
	}
}
