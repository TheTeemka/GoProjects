package services

import (
	"context"

	"github.com/TheTeemka/GoProjects/hw_6/models"
	"github.com/TheTeemka/GoProjects/hw_6/repository"
)

type ScheduleService struct {
	repo *repository.AttendanceRepository
}

func NewScheduleService(repo *repository.AttendanceRepository) *ScheduleService {
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
