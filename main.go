package main

import (
	"Jimbo8702/randomThoughts/cosmos/config"
	"Jimbo8702/randomThoughts/cosmos/internal/database"
	"fmt"
	"log"
	"os"
	"reflect"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal(err)
	}
	// sql db type
	db, err := database.New[database.Pool](config.BuildDSN(), os.Getenv("DATABASE_TYPE"))
	if err != nil {
		log.Fatal(err)
	}
	t := reflect.TypeOf(db.Pool)
	fmt.Printf("Type of db: %s\n", t)

	//mongo db type
	mdb, err := database.New[database.Pool](os.Getenv("MONGO_DB_URL"), "mongo")
	if err != nil {
		log.Fatal(err)
	}
	t2 := reflect.TypeOf(mdb.Pool)
	fmt.Printf("Type of db: %s\n", t2)
}