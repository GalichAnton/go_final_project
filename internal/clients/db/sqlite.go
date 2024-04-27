package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const fallback = ` CREATE TABLE scheduler (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            date TEXT,
            title TEXT,
            comment TEXT,
            repeat TEXT
        );

        CREATE INDEX idx_scheduler_date ON scheduler(date);
`

type Client interface {
	DB() *sql.DB
	Close() error
}

type sqliteClient struct {
	db *sql.DB
}

func New(filepath string) (Client, error) {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		db, err := sql.Open("sqlite3", filepath)
		if err != nil {
			log.Fatalf("Unable to open sqlite db file: %v", err)
		}

		_, err = db.Exec(fallback)
		if err != nil {
			return nil, err
		}

		return &sqliteClient{db: db}, nil

	} else if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatalf("Unable to open sqlite db file: %v", err)
	}

	return &sqliteClient{
		db: db,
	}, nil
}

func (s *sqliteClient) DB() *sql.DB {
	return s.db
}

func (s *sqliteClient) Close() error {
	if s.db != nil {
		s.db.Close()
	}

	return nil
}
