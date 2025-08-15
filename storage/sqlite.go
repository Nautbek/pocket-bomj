package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"pocket-bomj/src"

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
			health INTEGER NOT NULL DEFAULT 100,
			money REAL NOT NULL DEFAULT 0.00
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

func (s *Storage) UpdateBomj(id int64, health uint8, money float32) error {
	const op = "storage.sqlite.UpdateBomj"

	stmt, err := s.db.Prepare("UPDATE bomjs SET (health, money) = (?, ?) WHERE id = ?;")

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(health, money, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetBomj(id int64) (error, *src.Bomj) {
	const op = "storage.sqlite.GetBomj"

	stmt, err := s.db.Prepare("SELECT * FROM bomjs WHERE id = ?;")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err), nil
	}

	b := src.Bomj{}
	row := stmt.QueryRow(id).Scan(&b.Id, &b.Money, &b.Health)

	if errors.Is(err, sql.ErrNoRows) {
		return errors.New(op + fmt.Sprintf(" Bomj with id `%v` not found", id)), nil
	}

	fmt.Println(row)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err), nil
	}

	return nil, &b
}
