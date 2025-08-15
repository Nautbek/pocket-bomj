package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(StoragePath string) (*Storage, error) {
	const op = "storage.sqlite.NewStorage"

	db, err := sql.Open("sqlite3", StoragePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS bomjs (
			id INTEGER PRIMARY KEY,
			health INTEGER,
			money 
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateBomj(health uint8) (int64, error) {
	const op = "storage.sqlite.CreateBomj"

	stmt, err := s.db.Prepare("INSERT INTO bomjs (health) VALUES (?);")

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(health)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (s *Storage) UpdateBomj(health uint8, money float32) error {
	const op = "storage.sqlite.UpdateBomj"

	stmt, err := s.db.Prepare("UPDATE bomjs SET (health, money) VALUES (?, ?);")

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(health, money)
	if err != nil {
		return err
	}

	return nil
}
