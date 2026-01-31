package services

import (
	"context"
	"fmt"

	"github.com/TheTeemka/GoProjects/hw/models"
)

// IStudentRepository defines methods StudentsService depends on.
type IStudentRepository interface {
	CreateStudent(ctx context.Context, student *models.StudentEntity) error
	GetStudentByID(ctx context.Context, id int) (*models.StudentEntity, error)
	GetStudentsByGroupID(ctx context.Context, groupID int) ([]*models.StudentEntity, error)
	ListStudents(ctx context.Context, filter *models.StudentFilter) ([]*models.StudentEntity, error)
	UpdateStudent(ctx context.Context, student *models.StudentEntity) error
	DeleteStudent(ctx context.Context, id int) error
}

// StudentsService provides operations around students.
type StudentsService struct {
	repo IStudentRepository
}

func NewStudentsService(repo IStudentRepository) *StudentsService {
	return &StudentsService{repo: repo}
}

func (s *StudentsService) CreateStudent(ctx context.Context, req *models.CreateStudentRequest) error {
	return s.repo.CreateStudent(ctx, req.ToEntity())
}

func (s *StudentsService) GetStudentByID(ctx context.Context, id int) (*models.StudentDTO, error) {
	entity, err := s.repo.GetStudentByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("err in GetStudentByID: %w", err)
	}

	return entity.ToDTO(), nil
}

func (s *StudentsService) GetStudentsByGroupID(ctx context.Context, groupID int) ([]*models.StudentDTO, error) {
	entities, err := s.repo.GetStudentsByGroupID(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("err in GetStudentsByGroupID: %w", err)
	}

	dtos := make([]*models.StudentDTO, len(entities))
	for i, e := range entities {
		dtos[i] = e.ToDTO()
	}

	return dtos, nil
}

func (s *StudentsService) ListStudents(ctx context.Context, filter *models.StudentFilter) ([]*models.StudentDTO, error) {
	entities, err := s.repo.ListStudents(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("err in ListStudents: %w", err)
	}

	dtos := make([]*models.StudentDTO, len(entities))
	for i, e := range entities {
		dtos[i] = e.ToDTO()
	}

	return dtos, nil
}

func (s *StudentsService) UpdateStudent(ctx context.Context, id int, req *models.UpdateStudentRequest) error {
	entity, err := s.repo.GetStudentByID(ctx, id)
	if err != nil {
		return fmt.Errorf("err in UpdateStudent: %w", err)
	}
	entity.PatchFromRequest(req)
	return s.repo.UpdateStudent(ctx, entity)
}

func (s *StudentsService) DeleteStudent(ctx context.Context, id int) error {
	return s.repo.DeleteStudent(ctx, id)
}
