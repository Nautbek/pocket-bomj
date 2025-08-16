package storage

import (
	"context"
	"fmt"
	"pocket-bomj/src"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Storage struct {
	db      *gorm.DB
	context context.Context
}

func NewStorage(StoragePath string) (*Storage, error) {
	const op = "storage.sqlite.NewStorage"

	dbGorm, err := gorm.Open(sqlite.Open(StoragePath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = dbGorm.AutoMigrate(&src.Bomj{})
	if err != nil {
		return nil, err
	}

	return &Storage{
		db:      dbGorm,
		context: context.Background(),
	}, nil
}

func (s *Storage) CreateBomj(b *src.Bomj) error {
	return s.db.Create(b).Error
}

func (s *Storage) UpdateBomj(b *src.Bomj) error {
	return s.db.Updates(b).Error
}

func (s *Storage) GetBomj(bId int64) *src.Bomj {
	var bomj *src.Bomj
	s.db.First(&bomj, bId)

	return bomj
}
