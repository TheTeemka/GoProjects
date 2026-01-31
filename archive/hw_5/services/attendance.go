package services

import (
	"github.com/TheTeemka/GoProjects/hw_5/models"
	"github.com/TheTeemka/GoProjects/hw_5/repository"
)

type AttendanceService struct {
	repo *repository.AttendanceRepository
}

func NewAttendanceService(repo *repository.AttendanceRepository) *AttendanceService {
	return &AttendanceService{
		repo: repo,
	}
}

func (as *AttendanceService) CreateAttendance(req *models.CreateAttendanceDTO) error {
	entity := req.ToAttendanceEntity()
	return as.repo.CreateAttendance(entity)
}

func (as *AttendanceService) GetAllAttendanceByStudentID(studentID int) ([]*models.AttendanceDTO, error) {
	entities, err := as.repo.GetAllAttendanceByStudentID(studentID)
	if err != nil {
		return nil, err
	}
	dtos := make([]*models.AttendanceDTO, len(entities))
	for i, entity := range entities {
		dtos[i] = entity.ToAttendanceDTO()
	}
	return dtos, nil
}

func (as *AttendanceService) GetAllAttendanceBySubjectID(subjectID int) ([]*models.AttendanceDTO, error) {
	entities, err := as.repo.GetAllAttendanceBySubjectID(subjectID)
	if err != nil {
		return nil, err
	}
	dtos := make([]*models.AttendanceDTO, len(entities))
	for i, entity := range entities {
		dtos[i] = entity.ToAttendanceDTO()
	}
	return dtos, nil
}
