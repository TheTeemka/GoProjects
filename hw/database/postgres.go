package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func OpenConnection(dbString string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dbString)
	if err != nil {
		panic(err)
	}
	return conn
}
