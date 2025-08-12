package handlers

import (
	"pocket-bomj/repository"
	"pocket-bomj/src"
)

func NewBomj() (*src.Bomj, error) {
	bomj := src.Bomj{Health: 100}

	if err := repository.NewBomjRepository().Create(&bomj); err != nil {
		return nil, err
	}

	return &bomj, nil
}
