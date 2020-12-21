package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(dsn string) (*Repository, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, fmt.Errorf("Opening database failed: %v", err)
	}

	return &Repository{db: db}, nil
}
