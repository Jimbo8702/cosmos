package cosmos

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func buildRouter(debug bool) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	if debug {
		mux.Use(middleware.Logger)
	}
	mux.Use(middleware.Recoverer)	
	
	// //	add our session middleware
	// mux.Use(c.SessionLoad)

	return mux
}