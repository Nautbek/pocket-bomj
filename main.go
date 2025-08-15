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

	b.PlusMoney(190)

	log.Println(b)
}
