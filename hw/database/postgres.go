package database

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

func OpenConnection(dbString string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dbString)
	if err != nil {
		panic(err)
	}
	return conn
}

func PGXConnToSQLDB(conn *pgx.Conn) *sql.DB {
	sqlDB := sql.OpenDB(stdlib.GetConnector(*conn.Config()))

	return sqlDB
}
