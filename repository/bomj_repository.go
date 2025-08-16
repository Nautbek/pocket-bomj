package repository

import (
	"errors"
	"log"
	"pocket-bomj/src"
	"pocket-bomj/storage"
	"sync"
)

type BomjRepository struct {
	instance *src.Bomj
	once     sync.Once
	storage  storage.Storage
}

func (br *BomjRepository) Create(b *src.Bomj) error {
	return br.storage.CreateBomj(b)
}

func (br *BomjRepository) Update(b *src.Bomj) error {
	return br.storage.UpdateBomj(b)
}

func (br *BomjRepository) Get(id int64) (*src.Bomj, error) {
	bomj := br.storage.GetBomj(id)
	if bomj == nil {
		return nil, errors.New("bomj not found")
	}
	return br.storage.GetBomj(id), nil
}

var (
	instance *BomjRepository
	once     sync.Once
)

func NewBomjRepository() *BomjRepository {
	once.Do(func() {
		instance = &BomjRepository{}
		storage1, err := storage.NewStorage("storage/storage.db")
		if err != nil {
			log.Fatal(err)
		}
		instance.storage = *storage1
	})

	return instance
}
