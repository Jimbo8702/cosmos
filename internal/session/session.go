package session

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Jimbo8702/cosmos/config"
	"github.com/Jimbo8702/cosmos/internal/database"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName 	   string
	CookieDomain   string
	SessionType    string
	CookieSecure   string
	DBPool 		   database.Pool
}

func New(con *config.SessionConfig) *Session {
	return &Session{
		CookieLifetime: con.CookieLifetime,
		CookiePersist: con.CookiePersist,
		CookieName: con.CookieName,
		CookieDomain: con.CookieDomain,
		CookieSecure: con.CookieSecure,
		SessionType: con.SessionType,
	}
}

func(c *Session) Init(db *database.Database[database.Pool]) (*scs.SessionManager, error) {
	//defaults to false
	var persist, secure bool

	// how long should sessions last? 
	minutes, err := strconv.Atoi(c.CookieLifetime)
	if err != nil {
		minutes = 60
	}

	// should cookies persist?
	if strings.ToLower(c.CookiePersist) == "true" {
		persist = true
	} 

	// must cookies be secure? 
	if strings.ToLower(c.CookieSecure) == "true" {
		secure = true
	} 

	// create session
	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Secure = secure
	session.Cookie.Name = c.CookieName
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// which session store?
	switch strings.ToLower(c.SessionType) {
	case "redis":
		// if redis we would create a new redis instance with the database.New() then do type assertion, then add it to the session store
	case "postgres", "postgresql":
		db, ok := db.Pool.(*sql.DB)
		if !ok {
			return nil, errors.New("unable to determine database type for sessions")
		}
		session.Store = postgresstore.New(db)
	default:
		//defaults to cookie auth
	}
	return session, nil
}