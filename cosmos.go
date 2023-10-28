package main

import (
	"Jimbo8702/randomThoughts/cosmos/config"
	"Jimbo8702/randomThoughts/cosmos/internal/database"
	"Jimbo8702/randomThoughts/cosmos/internal/render"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

type DBOption[] struct {

}

type Application[db database.Pool] struct {
	//config values
	AppName 		string 
	Debug 			bool
	Version 		string
	RootPath 		string

	//external items
	ErrorLog 		*log.Logger
	InfoLog  		*log.Logger
	Routes 	 		*chi.Mux
	Session	 		*scs.SessionManager
	Engine 			render.Renderer
	Database 		database.Database[db]
}

func New[db database.Pool]() *Application[db] {
	return &Application[db]{}
}

func NewApp(rootPath string) error {
	var (
		app *Application[database.Pool]
		// err error
	)
	if err := godotenv.Load(rootPath + "/.env"); err != nil {
		return err
	}

	config := config.BuildConfig()
	//db type to app
	// switch strings.ToLower(config.DATABASE_TYPE) {
	// case "postgres", "postgresql":
	// 	app, err := InitSQLApp[*sql.DB](config.DATABASE_TYPE)
	// 	if err != nil {
	// 		return err
	// 	}
	// case "mongo":
	// 	app, err = InitMongoApp[database.Pool](config.DATABASE_TYPE)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	
	app.RootPath = config.RootPath
	// build db
	// build loggers
	// build session
	// build render 
	return nil
}

// func InitSQLApp[p *sql.DB](dbType string) (*Application[p], error) {
// 	db, err := database.NewSQL[*sql.DB](dbType, config.BuildDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return &Application[p]{
// 		Database: db,
// 	}, nil
// }

// func InitMongoApp[a database.Pool](dbType string) (*Application[a], error) {
// 	return &Application[a]{}, nil
// }

