package handlers

import (
	"pocket-bomj/repository"
	"pocket-bomj/src"
	"testing"
)

type MockBomjRepository struct {
	bomj *src.Bomj
	err  error
}

func (m *MockBomjRepository) Get(id int64) (*src.Bomj, error) {
	return m.bomj, m.err
}

func (m *MockBomjRepository) Update(b *src.Bomj) error {
	return m.err
}

func (m *MockBomjRepository) Create(b *src.Bomj) error {
	return m.err
}

func (m *MockBomjRepository) NewBomjRepository() repository.IBomjRepository {
	return m
}

func TestAddMoneyHandler(t *testing.T) {
	mockRepo := &MockBomjRepository{
		bomj: &src.Bomj{Id: 1, Money: 100},
	}

	originalRepo := repository.OBomjRepository
	repository.OBomjRepository = mockRepo.NewBomjRepository
	defer func() {
		repository.OBomjRepository = originalRepo
	}()

	err := AddMoneyHandler(1, 100)
	if err != nil {
		t.Error(err)
	}
}
