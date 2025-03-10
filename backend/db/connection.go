package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Database interface {
	Close(context.Context) error
	Ping(context.Context) error
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

var DBConn Database
