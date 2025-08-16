package main

import (
	"log"
	"pocket-bomj/handlers"
	"pocket-bomj/repository"
)

func main() {
	repository.NewBomjRepository()

	b, err := handlers.NewBomj()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(b)

	err = handlers.AddMoneyHandler(b.GetId(), float32(23190))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(b)

	bomj, err := handlers.GetBomj(b.GetId())
	if err != nil {
		log.Fatal(err)
	}

	log.Println(bomj)

	log.Println(handlers.GetBomj(123))
}
