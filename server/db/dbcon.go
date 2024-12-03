package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DBTX interface is compatible with pgxpool.Pool and pgx.Tx.
type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Ping(context.Context) error
}

// DB is the global database connection pool.
var DB DBTX

// DatabaseConnection initializes the database connection and returns a DBTX implementation.
func DatabaseConnection() (DBTX, error) {
	dbUri := os.Getenv("DB_URI")
	if dbUri == "" {
		log.Fatal("DB_URI environment variable is not set")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbUri)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	// Assign the pool to the global variable
	DB = pool
	return pool, nil
}

func GetDbCon() DBTX {
	return DB
}
