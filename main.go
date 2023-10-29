package main

import (
	"Jimbo8702/randomThoughts/cosmos/cosmos"
	"log"
	"os"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	app, err := cosmos.New(path)
	if err != nil {
		log.Fatal(err)
	}
	
	log.Fatal(app.ListenAndServe())
}