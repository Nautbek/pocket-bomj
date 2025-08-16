package handlers

import (
	"pocket-bomj/repository"
	"pocket-bomj/src"
)

func NewBomj() (*src.Bomj, error) {
	bomj := src.Bomj{}
	bomj.SetHealth(100)
	bomj.PlusMoney(100)

	if err := repository.OBomjRepository().Create(&bomj); err != nil {
		return nil, err
	}

	return &bomj, nil
}
