package repository

import (
	"context"
	"errors"

	"github.com/TheTeemka/GoProjects/hw/errs"
	"github.com/TheTeemka/GoProjects/hw/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type UserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *models.UserEntity) error {
	query := `INSERT INTO users ( email, role, password_hash) VALUES ($1, $2, $3)`

	_, err := ur.conn.Exec(ctx, query, user.Email, user.Role, user.PasswordHash)
	if err != nil {
		return ur.HandleSQLerr(err)
	}
	return nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.UserEntity, error) {
	query := `SELECT id, email, role, password_hash FROM users WHERE email=$1`
	row := ur.conn.QueryRow(ctx, query, email)

	var user models.UserEntity
	err := row.Scan(&user.ID, &user.Email, &user.Role, &user.PasswordHash)
	if err != nil {
		return nil, ur.HandleSQLerr(err)
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, user_id int) (*models.UserEntity, error) {
	query := `SELECT id, email, role, password_hash FROM users WHERE id=$1`
	row := ur.conn.QueryRow(ctx, query, user_id)

	var user models.UserEntity
	err := row.Scan(&user.ID, &user.Email, &user.Role, &user.PasswordHash)
	if err != nil {
		return nil, ur.HandleSQLerr(err)
	}

	return &user, nil
}

func (ur *UserRepository) HandleSQLerr(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		// 23505 is unique_violation
		if pgErr.Code == "23505" {
			return errs.ErrUserAlreadyExists
		}
	}

	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return errs.ErrUserAlreadyExists
		}
	}

	return err
}
