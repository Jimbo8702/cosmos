package session

import (
	"Jimbo8702/randomThoughts/cosmos/config"
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	DBPool 		   *sql.DB
}

func New(con *config.Config) (*scs.SessionManager, error) {
	sess := &Session{
		CookieLifetime: con.CookieLifetime,
		CookiePersist: con.CookiePersist,
		CookieName: con.CookieName,
		CookieDomain: con.CookieDomain,
		CookieSecure: con.CookieSecure,
		SessionType: con.SessionType,
	}
	return sess.init(), nil
}

func(c *Session) init() *scs.SessionManager {
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
		//
	case "postgres", "postgresql":
		session.Store = postgresstore.New(c.DBPool)
	default:
		//defaults to cookie auth
	}
	return session
}