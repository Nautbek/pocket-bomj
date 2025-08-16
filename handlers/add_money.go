package handlers

import (
	"errors"
	"pocket-bomj/repository"
)

func AddMoneyHandler(bId int64, money float32) error {
	if money <= 0 {
		return errors.New("money must be greater than zero")
	}

	bomjRepository := repository.NewBomjRepository()

	b, err := bomjRepository.Get(bId)

	if err != nil {
		return err
	}

	b.PlusMoney(money)

	if err := bomjRepository.Update(b); err != nil {
		return err
	}

	return nil
}
