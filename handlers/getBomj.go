package handlers

import (
	"pocket-bomj/repository"
	"pocket-bomj/src"
)

func GetBomj(id int64) (error, *src.Bomj) {
	return repository.NewBomjRepository().Get(id)
}
