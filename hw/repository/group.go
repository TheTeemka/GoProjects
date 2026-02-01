package repository

import (
	"context"

	"github.com/TheTeemka/GoProjects/hw/errs"
	"github.com/TheTeemka/GoProjects/hw/models"
	"github.com/jackc/pgx/v4"
)

type GroupRepository struct {
	conn *pgx.Conn
}

func NewGroupRepository(conn *pgx.Conn) *GroupRepository {
	return &GroupRepository{conn: conn}
}

func (r *GroupRepository) CreateGroup(ctx context.Context, group *models.GroupEntity) (int, error) {
	query := `
        INSERT INTO groups (name, enrollment_year)
        VALUES ($1, $2)
		RETURNING id`

	row := r.conn.QueryRow(ctx, query, group.Name, group.EnrollmentYear)

	var id int
	err := row.Scan(&id)

	return id, err
}

func (r *GroupRepository) GetGroupByID(ctx context.Context, id int) (*models.GroupEntity, error) {
	query := `
        SELECT id, name, enrollment_year, created_at, updated_at
        FROM groups
        WHERE id = $1`
	row := r.conn.QueryRow(ctx, query, id)

	var group models.GroupEntity
	err := row.Scan(
		&group.ID,
		&group.Name,
		&group.EnrollmentYear,
		&group.CreatedAt,
		&group.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.ErrGroupNotFound
		}
		return nil, err
	}

	return &group, nil
}
