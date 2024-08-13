package main

import (
	"log"

	"github.com/itsmoirob/ecom-auth/cmd/api"
	"github.com/itsmoirob/ecom-auth/db"
)

func main() {
	store, err := db.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":28940", store)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
