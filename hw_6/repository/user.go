package repository

import (
	"context"

	"github.com/TheTeemka/GoProjects/hw_5/models"
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
	query := `INSERT INTO users ( email, password_hash) VALUES ($1, $2)`

	_, err := ur.conn.Exec(ctx, query, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.UserEntity, error) {
	query := `SELECT id, email, password_hash FROM users WHERE email=$1`

	row := ur.conn.QueryRow(ctx, query, email)

	var user models.UserEntity
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {

		return nil, err
	}
	return &user, nil
}
