package repository

import (
	"context"
	"errors"
	"time"

	"github.com/TheTeemka/GoProjects/hw/errs"
	"github.com/jackc/pgx/v4"
)

type TokenRepository struct {
	conn *pgx.Conn
}

func NewTokenRepository(conn *pgx.Conn) *TokenRepository {
	return &TokenRepository{
		conn: conn,
	}
}

func (tr *TokenRepository) CreateToken(ctx context.Context, userID int, token string, expiryDate time.Time, tokenType string) error {
	query := `INSERT INTO tokens (user_id, token, expires_at, type) VALUES ($1, $2, $3, $4)`

	_, err := tr.conn.Exec(ctx, query, userID, token, expiryDate, tokenType)
	if err != nil {
		return err
	}
	return nil
}

func (tr *TokenRepository) TokenExists(ctx context.Context, token string, tokenType string) (int, error) {
	query := `SELECT user_id FROM tokens WHERE token=$1 AND type=$2 AND expires_at > NOW()`
	row := tr.conn.QueryRow(ctx, query, token, tokenType)

	var user_id int
	err := row.Scan(&user_id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, errs.ErrTokenNotFound
		}
		return 0, err
	}

	return user_id, nil
}
