package repository

import (
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
	id, err := br.storage.CreateBomj(b.GetHealth())
	if err != nil {
		return err
	}

	b.SetId(id)
	return nil
}

func (br *BomjRepository) Update(b src.BomjInterface) error {
	return br.storage.UpdateBomj(b.GetId(), b.GetHealth(), b.GetMoney())
}

func (br *BomjRepository) Get(id int64) (error, *src.Bomj) {
	return br.storage.GetBomj(id)
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
