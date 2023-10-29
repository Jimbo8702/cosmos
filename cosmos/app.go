package cosmos

import (
	"Jimbo8702/randomThoughts/cosmos/config"
	"Jimbo8702/randomThoughts/cosmos/internal/database"
	"Jimbo8702/randomThoughts/cosmos/internal/logger"
	"Jimbo8702/randomThoughts/cosmos/internal/render"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
)

const version = "0.1.0"

type Application struct {
	// Application configuration
	AppName 		string 
	Debug 			bool
	Version 		string
	RootPath 		string
	Port 			string

	// Internal 
	// 
	// Application error log
	ErrorLog 		logger.Logger
	//
	// Application info log
	InfoLog  		logger.Logger
	//
	// Application render engine
	Engine 			render.Renderer
	//
	// Application database connection
	Database 		*database.Database[database.Pool]
	//
	// THIRD PARTY DEPS
	//
	// Chi mux for routing
	Router	 		*chi.Mux
	//
	// scs session manager for handling secure sessions
	Session	 		*scs.SessionManager
	// might make this an interface to allow for different kinds of session management
	// idk tho thats a lot
}

// Build a new app with config vars loaded from env
func NewApp(c *config.Config) *Application {
	return &Application{
		RootPath: c.RootPath,
		AppName: c.AppName,
		Debug: c.Debug,
		Version: version,
		Port: c.Port,
	}
}

// Starts the http server 
func (a *Application) ListenAndServe() error {
	w := a.ErrorLog.GetLogWriter()
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", a.Port),
		ErrorLog: log.New(w, "", 0),
		Handler: a.Router,
		IdleTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}
	// figure out how to capture errors with the defer
	// defer closing database connection
	defer a.CloseDB()

	// log port and return listen and serve from the server 
	a.InfoLog.Log("Listening on port %s", a.Port)
	return srv.ListenAndServe()
}

func  (a *Application) CloseDB() error {
	var err error
	switch a.Database.Type {

	// Handle closing the sql connection
	case "postgres", "postgresql":
		db, ok := a.Database.Pool.(*sql.DB)
		if !ok {
			return errors.New("unable to cast to sql client with sql db type")
		}
		err = db.Close()

	// Handle closing the mongo connection
	case "mongo":
		db, ok := a.Database.Pool.(*mongo.Client)
		if !ok {
			return errors.New("unable to cast to mongo client with mongo db type")
		}
		err = db.Disconnect(context.Background())

	// type unsupported, technically should never get here 
	default:
		err = errors.New("couldn't find db type while closing")
	}
	return err
}