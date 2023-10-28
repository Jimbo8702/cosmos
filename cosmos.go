package main

import (
	"Jimbo8702/randomThoughts/cosmos/config"
	"Jimbo8702/randomThoughts/cosmos/internal/database"
	"Jimbo8702/randomThoughts/cosmos/internal/render"
	"Jimbo8702/randomThoughts/cosmos/internal/session"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

const version = "0.1.0"

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

func New[db database.Pool](c *config.Config) *Application[db] {
	return &Application[db]{
		RootPath: c.RootPath,
		AppName: c.AppName,
		Debug: c.Debug,
		Version: version,
	}
}

func NewApp(rootPath string) error {
	if err := godotenv.Load(rootPath + "/.env"); err != nil {
		return err
	}
	c := config.BuildConfig()
	app := New[database.Pool](c)

	// DATABASE SETUP
	db, err := database.New[database.Pool](config.BuildDSN(), c.DATABASE_TYPE)
	if err != nil {
		return err
	}
	app.Database = *db

	// SESSION SETUP
	sess := session.New(&c.SessionConfig)
	manager, err := sess.Init(app.Database)
	if err != nil {
		return err
	}
	app.Session = manager
	
	// RENDER ENGINE SETUP
	renderer, err := render.New(app.Session, c)
	if err != nil {
		return err
	}
	app.Engine = renderer

	// build loggers

	return nil
}






