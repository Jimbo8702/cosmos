package cosmos

import (
	"Jimbo8702/randomThoughts/cosmos/config"
	"Jimbo8702/randomThoughts/cosmos/internal/database"
	"Jimbo8702/randomThoughts/cosmos/internal/logger"
	"Jimbo8702/randomThoughts/cosmos/internal/render"
	"Jimbo8702/randomThoughts/cosmos/internal/session"

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
	ErrorLog 		logger.Logger
	InfoLog  		logger.Logger
	Routes 	 		*chi.Mux
	Session	 		*scs.SessionManager
	Engine 			render.Renderer
	Database 		database.Database[db]
}

func NewApp[db database.Pool](c *config.Config) *Application[db] {
	return &Application[db]{
		RootPath: c.RootPath,
		AppName: c.AppName,
		Debug: c.Debug,
		Version: version,
	}
}

func New(rootPath string) (*Application[database.Pool], error) {
	// load  env
	if err := godotenv.Load(rootPath + "/.env"); err != nil {
		return nil, err
	}
	// load config
	c := config.Load()
	// build new app with config
	app := NewApp[database.Pool](c)

	// DATABASE SETUP 
	db, err := database.New[database.Pool](config.BuildDSN(), c.DATABASE_TYPE)
	if err != nil {
		return nil, err
	}
	app.Database = *db

	// SESSION SETUP
	sess := session.New(&c.SessionConfig)
	manager, err := sess.Init(app.Database)
	if err != nil {
		return nil, err
	}
	app.Session = manager
	
	// RENDER ENGINE SETUP
	renderer, err := render.New(app.Session, c)
	if err != nil {
		return nil, err
	}
	app.Engine = renderer

	// LOGGER SETUP
	app.ErrorLog = logger.NewErrorLog()
	app.InfoLog = logger.NewInfoLog()

	return app, nil
}
