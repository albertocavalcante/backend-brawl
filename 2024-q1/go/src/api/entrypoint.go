package api

import (
	"backend-brawl/src/db"
	"fmt"
)

func Execute() {
	fmt.Println("Executing...")

	store := &db.Store{}
	store.InitDatabase("file:sqlite.db?cache=shared&mode=rwc")
	store.CreateSchemas()
	store.PopulateDatabase()

	server := &Server{
		Port:  9003,
		Store: store,
	}

	server.ListenAndServe()
}
