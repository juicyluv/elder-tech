package main

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New("file://migrations", "postgres://postgres:postgres@localhost:5432/elder?sslmode=disable")
	if err != nil {
		log.Fatalf("creating migration: %v", err)
	}
	if err = m.Drop(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("applying migrations: %v", err)
	}

	log.Println("DROP done.")
}
