package repository

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	Db *sql.DB
}

func New(dsn string) (*Repository, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, fmt.Errorf("Opening database failed: %v", err)
	}

	return &Repository{Db: db}, nil
}
