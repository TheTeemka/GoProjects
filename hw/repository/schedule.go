package repository

import (
	"context"

	"github.com/TheTeemka/GoProjects/hw/models"

	"github.com/jackc/pgx/v4"
)

type ScheduleRepository struct {
	conn *pgx.Conn
}

func NewScheduleRepository(conn *pgx.Conn) *AttendanceRepository {
	return &AttendanceRepository{
		conn: conn,
	}
}

func (sc *AttendanceRepository) GetForStudent(ctx context.Context, studentID int) ([]models.Schedule, error) {
	query := `
        SELECT schedules.id, schedules.subject, schedules.day_of_week, schedules.time, schedules.group_id 
        FROM students 
        JOIN schedules ON schedules.group_id = students.group_id 
        WHERE students.id = $1`

	rows, err := sc.conn.Query(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []models.Schedule
	for rows.Next() {
		var schedule models.Schedule
		err := rows.Scan(
			&schedule.ID,
			&schedule.Subject,
			&schedule.DayOfWeek,
			&schedule.Time,
			&schedule.GroupID,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

func (sc *AttendanceRepository) GetAll(ctx context.Context) ([]models.Schedule, error) {
	rows, err := sc.conn.Query(
		ctx,
		"SELECT id, subject, day_of_week, time, group_id FROM schedules",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []models.Schedule
	for rows.Next() {
		var schedule models.Schedule
		err := rows.Scan(
			&schedule.ID,
			&schedule.Subject,
			&schedule.DayOfWeek,
			&schedule.Time,
			&schedule.GroupID,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

func (sc *AttendanceRepository) GetForGroup(ctx context.Context, groupID int) ([]models.Schedule, error) {
	rows, err := sc.conn.Query(
		ctx,
		"SELECT id, subject, day_of_week, time, group_id FROM schedules WHERE group_id = $1",
		groupID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []models.Schedule
	for rows.Next() {
		var schedule models.Schedule
		err := rows.Scan(
			&schedule.ID,
			&schedule.Subject,
			&schedule.DayOfWeek,
			&schedule.Time,
			&schedule.GroupID,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}
