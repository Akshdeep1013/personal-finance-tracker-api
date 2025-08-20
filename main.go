package main

import (
	"fmt"
	"log"
	"personal-finance-tracker-api/repository"
)

func main() {
	repo, err := repository.NewRepository()
	if err != nil {
		log.Fatal("\nUnable to initalize the database: ", err)
	}
	fmt.Printf("\n Database has been initialized!!\n")
	fmt.Printf("\n connection details : %+v", *repo)
}
