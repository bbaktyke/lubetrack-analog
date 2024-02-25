package repository

import "github.com/jmoiron/sqlx"

type Database interface {
}

type OilAnalysis struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) (Database, error) {
	return OilAnalysis{
		db: db,
	}, nil
}
