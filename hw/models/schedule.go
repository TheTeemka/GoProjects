package models

import "time"

const timeFormat = "15:04"

type ScheduleDTO struct {
	ID        int    `json:"id"`
	Subject   string `json:"subject"`
	DayOfWeek int    `json:"day_of_week"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	GroupID   int    `json:"group_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ScheduleEntity struct {
	ID        int
	Subject   string
	DayOfWeek int
	StartTime time.Time
	EndTime   time.Time
	GroupID   int
	CreatedAt string
	UpdatedAt string
}

func (e *ScheduleEntity) ToDTO() *ScheduleDTO {
	return &ScheduleDTO{
		ID:        e.ID,
		Subject:   e.Subject,
		DayOfWeek: e.DayOfWeek,
		StartTime: e.StartTime.Format(timeFormat),
		EndTime:   e.EndTime.Format(timeFormat),
		GroupID:   e.GroupID,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

type CreateScheduleRequest struct {
	Subject   string `json:"subject"`
	DayOfWeek int    `json:"day_of_week"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	GroupID   int    `json:"group_id"`
}

func (r *CreateScheduleRequest) ToEntity() (*ScheduleEntity, error) {
	parsedStartTime, err := time.Parse(timeFormat, r.StartTime)
	if err != nil {
		return nil, err
	}

	parsedEndTime, err := time.Parse(timeFormat, r.EndTime)
	if err != nil {
		return nil, err
	}

	return &ScheduleEntity{
		Subject:   r.Subject,
		DayOfWeek: r.DayOfWeek,
		StartTime: parsedStartTime,
		EndTime:   parsedEndTime,
		GroupID:   r.GroupID,
	}, nil
}
