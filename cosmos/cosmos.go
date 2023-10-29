package cosmos

import (
	"github.com/Jimbo8702/cosmos/config"
	"github.com/Jimbo8702/cosmos/internal/database"
	"github.com/Jimbo8702/cosmos/internal/logger"
	"github.com/Jimbo8702/cosmos/internal/render"
	"github.com/Jimbo8702/cosmos/internal/session"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func New(rootPath string) (*Application, error) {
	var err error
	// load  env
	if err = godotenv.Load(rootPath + "/.env"); err != nil {
		return nil, err
	}

	// load config
	c := config.Load()

	// build new app with config
	app := NewApp(c)

	// DATABASE SETUP 
	app.Database, err = database.New[database.Pool](config.BuildDSN(), c.DATABASE_TYPE)
	if err != nil {
		return nil, err
	}

	// SESSION SETUP
	sess := session.New(&c.SessionConfig)
	app.Session, err = sess.Init(app.Database)
	if err != nil {
		return nil, err
	}

	// if no render is present, then it just becomes an rest api (no view handling)
	if c.Renderer != "" {
		// RENDER ENGINE SETUP
		app.Engine, err = render.New(app.Session, c)
		if err != nil {
			return nil, err
		}
	}
	
	// LOGGER SETUP
	app.ErrorLog = logger.New(logger.ERROR)
	app.InfoLog = logger.New(logger.INFO)

	// ROUTER SETUP
	app.Router = buildRouter(c.Debug).(*chi.Mux)

	return app, nil
}
