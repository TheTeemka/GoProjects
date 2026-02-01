package services

import (
	"context"
	"fmt"

	"github.com/TheTeemka/GoProjects/hw/models"
)

// IGroupRepository defines repository methods GroupService depends on.
type IGroupRepository interface {
	CreateGroup(ctx context.Context, group *models.GroupEntity) (int, error)
	GetGroupByID(ctx context.Context, id int) (*models.GroupEntity, error)
}

// GroupService provides operations around groups.
type GroupService struct {
	repo IGroupRepository
}

func NewGroupService(repo IGroupRepository) *GroupService {
	return &GroupService{repo: repo}
}

func (s *GroupService) CreateGroup(ctx context.Context, req *models.CreateGroupRequest) (int, error) {
	return s.repo.CreateGroup(ctx, req.ToEntity())
}

func (s *GroupService) GetGroupByID(ctx context.Context, id int) (*models.GroupDTO, error) {
	entity, err := s.repo.GetGroupByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("err in GetGroupByID: %w", err)
	}
	return entity.ToDTO(), nil
}
