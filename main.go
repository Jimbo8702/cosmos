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
	
	app.ErrorLog.Log("this is an example error", nil)
	app.InfoLog.Log("this is an example Info", nil)

	//an example on how you can use different loggers 
	// standardErrorLog := &logger.StdLogger{}
	// standardErrorLog.SetLevel(logger.ERROR)
	// app.ErrorLog = standardErrorLog
	// app.ErrorLog.Log("this is now using the format package from the standard libaray", nil)

	// standardInfoLog := &logger.StdLogger{}
	// standardInfoLog.SetLevel(logger.INFO)
	// app.InfoLog = standardInfoLog
	// app.InfoLog.Log("this is now using the format package from standard libaray", nil)
}