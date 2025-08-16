package storage

import (
	"context"
	"database/sql"
	"fmt"
	"pocket-bomj/src"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Storage struct {
	db        *sql.DB
	dbGorm    *gorm.DB
	gormCntxt context.Context
}

func NewStorage(StoragePath string) (*Storage, error) {
	const op = "storage.sqlite.NewStorage"

	db, err := sql.Open("sqlite3", StoragePath)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	dbGorm, err := gorm.Open(sqlite.Open(StoragePath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
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

	return &Storage{
		db:        db,
		dbGorm:    dbGorm,
		gormCntxt: context.Background(),
	}, nil
}

func (s *Storage) CreateBomj(b *src.Bomj) error {
	return s.dbGorm.Create(b).Error
}

func (s *Storage) UpdateBomj(b *src.Bomj) error {
	return s.dbGorm.Updates(b).Error
}

func (s *Storage) GetBomj(bId int64) *src.Bomj {
	var bomj *src.Bomj
	s.dbGorm.First(&bomj, bId)

	return bomj
}
