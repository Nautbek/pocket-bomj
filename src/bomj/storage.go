package bomj

import (
	"sync"
)

// Storage представляет хранилище бомжей
type Storage struct {
	bomjs map[int64]*Bomj
	mutex sync.RWMutex
}

// NewStorage создает новое хранилище
func NewStorage() *Storage {
	return &Storage{
		bomjs: make(map[int64]*Bomj),
	}
}

// SaveBomj сохраняет бомжа
func (s *Storage) SaveBomj(bomj *Bomj) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.bomjs[bomj.UserID] = bomj
}

// GetBomj возвращает бомжа по ID пользователя
func (s *Storage) GetBomj(userID int64) *Bomj {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.bomjs[userID]
}

// HasBomj проверяет, есть ли у пользователя бомж
func (s *Storage) HasBomj(userID int64) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	_, exists := s.bomjs[userID]
	return exists
}

// GetAllBomjs возвращает всех бомжей (для админки)
func (s *Storage) GetAllBomjs() []*Bomj {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	bomjs := make([]*Bomj, 0, len(s.bomjs))
	for _, bomj := range s.bomjs {
		bomjs = append(bomjs, bomj)
	}
	return bomjs
}
