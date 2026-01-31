package repository

import (
	"context"
	"fmt"

	"github.com/TheTeemka/GoProjects/hw/models"

	"github.com/jackc/pgx/v4"
)

type AttendanceRepository struct {
	conn *pgx.Conn
}

func NewAttendanceRepository(conn *pgx.Conn) *AttendanceRepository {
	return &AttendanceRepository{
		conn: conn,
	}
}

func (ar *AttendanceRepository) CreateAttendance(attendance *models.AttendanceEntity) error {
	query := `
        INSERT INTO attendance (student_id, subject_id, visit_date, visited)
        VALUES ($1, $2, $3, $4)`

	_, err := ar.conn.Exec(context.Background(), query, attendance.StudentID, attendance.SubjectID, attendance.VisitDate, attendance.Visited)
	if err != nil {
		return fmt.Errorf("error in CreateAttendance: %v", err)
	}

	return nil
}

func (ar *AttendanceRepository) GetAllAttendanceByStudentID(studentID int) ([]*models.AttendanceEntity, error) {
	query := `
        SELECT id, student_id, subject_id, visit_date, visited
        FROM attendance
        WHERE student_id = $1`

	rows, err := ar.conn.Query(context.Background(), query, studentID)
	if err != nil {
		return nil, fmt.Errorf("error in GetAttendanceByStudentID: %v", err)
	}
	defer rows.Close()

	var attendances []*models.AttendanceEntity
	for rows.Next() {
		var attendance models.AttendanceEntity
		err := rows.Scan(&attendance.ID, &attendance.StudentID, &attendance.SubjectID, &attendance.VisitDate, &attendance.Visited)
		if err != nil {
			return nil, fmt.Errorf("error scanning row in GetAttendanceByStudentID: %v", err)
		}
		attendances = append(attendances, &attendance)
	}

	return attendances, nil
}

func (ar *AttendanceRepository) GetAllAttendanceBySubjectID(subjectID int) ([]*models.AttendanceEntity, error) {
	query := `
        SELECT id, student_id, subject_id, visit_date, visited
        FROM attendance
        WHERE subject_id = $1`

	rows, err := ar.conn.Query(context.Background(), query, subjectID)
	if err != nil {
		return nil, fmt.Errorf("error in GetAttendanceBySubjectID: %v", err)
	}
	defer rows.Close()

	var attendances []*models.AttendanceEntity
	for rows.Next() {
		var attendance models.AttendanceEntity
		err := rows.Scan(&attendance.ID, &attendance.StudentID, &attendance.SubjectID, &attendance.VisitDate, &attendance.Visited)
		if err != nil {
			return nil, fmt.Errorf("error scanning row in GetAttendanceBySubjectID: %v", err)
		}
		attendances = append(attendances, &attendance)
	}

	return attendances, nil
}
