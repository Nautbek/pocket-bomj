package handlers

import (
	"pocket-bomj/repository"
	"pocket-bomj/src"
)

func GetBomj(id int64) (*src.Bomj, error) {
	return repository.OBomjRepository().Get(id)
}
