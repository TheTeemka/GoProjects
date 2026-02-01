package services

import (
	"context"

	"github.com/TheTeemka/GoProjects/hw/models"
	"github.com/TheTeemka/GoProjects/hw/repository"
)

type ScheduleService struct {
	repo repository.ScheduleRepo
}

func NewScheduleService(repo repository.ScheduleRepo) *ScheduleService {
	return &ScheduleService{
		repo: repo,
	}
}

func (ss *ScheduleService) GetForStudent(ctx context.Context, studentID int) ([]*models.ScheduleDTO, error) {
	entities, err := ss.repo.GetForStudent(ctx, studentID)
	if err != nil {
		return nil, err
	}
	var out []*models.ScheduleDTO
	for _, e := range entities {
		out = append(out, e.ToDTO())
	}
	return out, nil
}

func (ss *ScheduleService) GetAll(ctx context.Context) ([]*models.ScheduleDTO, error) {
	entities, err := ss.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var out []*models.ScheduleDTO
	for _, e := range entities {
		out = append(out, e.ToDTO())
	}
	return out, nil
}

func (ss *ScheduleService) GetForGroup(ctx context.Context, groupID int) ([]*models.ScheduleDTO, error) {
	entities, err := ss.repo.GetForGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}
	var out []*models.ScheduleDTO
	for _, e := range entities {
		out = append(out, e.ToDTO())
	}
	return out, nil
}

func (ss *ScheduleService) CreateSchedule(ctx context.Context, req *models.CreateScheduleRequest) error {
	ent, err := req.ToEntity()
	if err != nil {
		return err
	}

	err = ss.repo.CreateSchedule(ctx, ent)
	if err != nil {
		return err
	}
	return nil
}

func (ss *ScheduleService) DeleteSchedule(ctx context.Context, id int) error {
	return ss.repo.DeleteSchedule(ctx, id)
}
