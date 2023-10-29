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
	//config values
	AppName 		string 
	Debug 			bool
	Version 		string
	RootPath 		string
	Port 			string

	//internal items
	ErrorLog 		logger.Logger
	InfoLog  		logger.Logger
	Engine 			render.Renderer
	Database 		*database.Database[database.Pool]

	//third party items
	Router	 		*chi.Mux
	Session	 		*scs.SessionManager
}

func NewApp(c *config.Config) *Application {
	return &Application{
		RootPath: c.RootPath,
		AppName: c.AppName,
		Debug: c.Debug,
		Version: version,
		Port: c.Port,
	}
}

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
	defer a.CloseDB()

	a.InfoLog.Log("Listening on port %s", a.Port)
	return srv.ListenAndServe()
}

func  (a *Application) CloseDB() error {
	var err error
	switch a.Database.Type {
	case "postgres", "postgresql":
		db, ok := a.Database.Pool.(*sql.DB)
		if !ok {
			return errors.New("unable to cast to sql client with sql db type")
		}
		err = db.Close()
	case "mongo":
		db, ok := a.Database.Pool.(*mongo.Client)
		if !ok {
			return errors.New("unable to cast to mongo client with mongo db type")
		}
		err = db.Disconnect(context.Background())
	default:
		err = errors.New("couldn't find db type while closing")
	}
	return err
}