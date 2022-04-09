package store

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	config *Config
	db     *sql.DB
}

func NewStore(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("mysql", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	log.Println(err)

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	log.Println(s.db)

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
