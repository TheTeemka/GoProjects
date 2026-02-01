package repository

import (
	"context"

	"github.com/TheTeemka/GoProjects/hw/models"

	"github.com/jackc/pgx/v4"
)

type ScheduleRepository struct {
	conn *pgx.Conn
}

// ScheduleRepo is the interface describing schedule repository behavior.
// It is implemented by *ScheduleRepository and used by services to allow easier testing.
type ScheduleRepo interface {
	GetForStudent(ctx context.Context, studentID int) ([]models.ScheduleEntity, error)
	GetAll(ctx context.Context) ([]models.ScheduleEntity, error)
	GetForGroup(ctx context.Context, groupID int) ([]models.ScheduleEntity, error)
	CreateSchedule(ctx context.Context, s *models.ScheduleEntity) error
	DeleteSchedule(ctx context.Context, id int) error
}

func NewScheduleRepository(conn *pgx.Conn) *ScheduleRepository {
	return &ScheduleRepository{
		conn: conn,
	}
}

func (sr *ScheduleRepository) GetForStudent(ctx context.Context, studentID int) ([]models.ScheduleEntity, error) {
	query := `
        SELECT schedules.id, schedules.subject, schedules.day_of_week, schedules.start_time, schedules.end_time, schedules.group_id 
        FROM students 
        JOIN schedules ON schedules.group_id = students.group_id 
        WHERE students.id = $1`

	rows, err := sr.conn.Query(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []models.ScheduleEntity
	for rows.Next() {
		var schedule models.ScheduleEntity
		err := rows.Scan(
			&schedule.ID,
			&schedule.Subject,
			&schedule.DayOfWeek,
			&schedule.StartTime,
			&schedule.EndTime,
			&schedule.GroupID,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

func (sr *ScheduleRepository) GetAll(ctx context.Context) ([]models.ScheduleEntity, error) {
	rows, err := sr.conn.Query(
		ctx,
		"SELECT id, subject, day_of_week, start_time, end_time, group_id FROM schedules",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []models.ScheduleEntity
	for rows.Next() {
		var schedule models.ScheduleEntity
		err := rows.Scan(
			&schedule.ID,
			&schedule.Subject,
			&schedule.DayOfWeek,
			&schedule.StartTime,
			&schedule.EndTime,
			&schedule.GroupID,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

func (sr *ScheduleRepository) GetForGroup(ctx context.Context, groupID int) ([]models.ScheduleEntity, error) {
	rows, err := sr.conn.Query(
		ctx,
		"SELECT id, subject, day_of_week, start_time, end_time, group_id FROM schedules WHERE group_id = $1",
		groupID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	schedules := make([]models.ScheduleEntity, 0)
	for rows.Next() {
		var schedule models.ScheduleEntity
		err := rows.Scan(
			&schedule.ID,
			&schedule.Subject,
			&schedule.DayOfWeek,
			&schedule.StartTime,
			&schedule.EndTime,
			&schedule.GroupID,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (sr *ScheduleRepository) CreateSchedule(ctx context.Context, s *models.ScheduleEntity) error {
	query := `
		INSERT INTO schedules (group_id, subject, day_of_week, start_time, end_time)
		VALUES ($1, $2, $3, $4, $5)`

	_, err := sr.conn.Exec(ctx, query, s.GroupID, s.Subject, s.DayOfWeek, s.StartTime, s.EndTime)
	if err != nil {
		return err
	}

	return nil
}

func (sr *ScheduleRepository) DeleteSchedule(ctx context.Context, id int) error {
	query := `
		DELETE FROM schedules
		WHERE id = $1`

	_, err := sr.conn.Exec(ctx, query, id)
	return err
}
