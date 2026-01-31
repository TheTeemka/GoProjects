package service

import (
	"context"
	"x/models"
	"x/repository"
)

type ScheduleService struct {
	repo *repository.ScheduleRepository
}

func NewScheduleService(repo *repository.ScheduleRepository) *ScheduleService {
	return &ScheduleService{
		repo: repo,
	}
}

func (ss *ScheduleService) GetForStudent(ctx context.Context, studentID int) ([]models.Schedule, error) {
	return ss.repo.GetForStudent(ctx, studentID)
}

func (ss *ScheduleService) GetAll(ctx context.Context) ([]models.Schedule, error) {
	return ss.repo.GetAll(ctx)
}

func (ss *ScheduleService) GetForGroup(ctx context.Context, groupID int) ([]models.Schedule, error) {
	return ss.repo.GetForGroup(ctx, groupID)
}
