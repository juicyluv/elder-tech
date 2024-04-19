package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func Init(pool *pgxpool.Pool) {
	db = pool
}
