package store

import (
	"database/sql"
	"https://github.com/jackc/pgx"
)

// Store ...
type Store struct {
	config *Config
	db     *sql.DB
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// New...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}
