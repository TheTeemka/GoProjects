package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/TheTeemka/GoProjects/hw/errs"
	"github.com/TheTeemka/GoProjects/hw/models"
	"github.com/jackc/pgx/v4"
)

type StudentRepository struct {
	conn *pgx.Conn
}

func NewStudentRepository(conn *pgx.Conn) *StudentRepository {
	return &StudentRepository{
		conn: conn,
	}
}

func (r *StudentRepository) CreateStudent(ctx context.Context, student *models.StudentEntity) error {
	query := `
		INSERT INTO students (name, email, birthday, group_id) 
		VALUES ($1, $2, $3, $4)`

	_, err := r.conn.Exec(
		ctx,
		query,
		student.Name,
		student.Email,
		student.Birthday,
		student.GroupID,
	)

	return err
}

func (r *StudentRepository) GetStudentByID(ctx context.Context, id int) (*models.StudentEntity, error) {
	query := `
		SELECT id, name, email, birthday, group_id, created_at, updated_at
		FROM students
		WHERE id = $1`

	row := r.conn.QueryRow(ctx, query, id)

	var student models.StudentEntity
	err := row.Scan(
		&student.ID,
		&student.Name,
		&student.Email,
		&student.Birthday,
		&student.GroupID,
		&student.CreatedAt,
		&student.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.ErrStudentNotFound
		}
		return nil, err
	}

	return &student, nil
}

func (r *StudentRepository) GetStudentsByGroupID(ctx context.Context, groupID int) ([]*models.StudentEntity, error) {
	query := `
		SELECT id, name, email, birthday, group_id, created_at, updated_at
		FROM students
		WHERE group_id = $1`

	rows, err := r.conn.Query(ctx, query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*models.StudentEntity
	for rows.Next() {
		var student models.StudentEntity
		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Email,
			&student.Birthday,
			&student.GroupID,
			&student.CreatedAt,
			&student.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	return students, nil
}

func (r *StudentRepository) ListStudents(ctx context.Context, filter *models.StudentFilter) ([]*models.StudentEntity, error) {
	query := squirrel.Select("id", "name", "email", "birthday", "group_id", "created_at", "updated_at").From("students")

	query = FilterToSQL(query, filter)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*models.StudentEntity
	for rows.Next() {
		var student models.StudentEntity
		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Email,
			&student.Birthday,
			&student.GroupID,
			&student.CreatedAt,
			&student.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	return students, nil
}

func (r *StudentRepository) UpdateStudent(ctx context.Context, student *models.StudentEntity) error {
	query := `
		UPDATE students
		SET name = $1, email = $2, birthday = $3, group_id = $4, updated_at = NOW()
		WHERE id = $5`

	_, err := r.conn.Exec(
		ctx,
		query,
		student.Name,
		student.Email,
		student.Birthday,
		student.GroupID,
		student.ID,
	)

	return err
}

func (r *StudentRepository) DeleteStudent(ctx context.Context, id int) error {
	query := `
		DELETE FROM students
		WHERE id = $1`

	_, err := r.conn.Exec(ctx, query, id)
	return err
}

func FilterToSQL(builder squirrel.SelectBuilder, filter *models.StudentFilter) squirrel.SelectBuilder {
	if filter.GroupID != nil {
		builder = builder.Where(squirrel.Eq{"group_id": *filter.GroupID})
	}

	if filter.Name != nil {
		builder = builder.Where(squirrel.Like{"name": "%" + *filter.Name + "%"})
	}

	if filter.Email != nil {
		builder = builder.Where(squirrel.Like{"email": "%" + *filter.Email + "%"})
	}

	if filter.Limit != 0 {
		builder = builder.Limit(filter.Limit)
	}

	if filter.Offset != 0 {
		builder = builder.Offset(filter.Offset)
	}

	return builder
}
